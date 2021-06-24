package utils

import (
	"github.com/spacerouter/docker_api/models"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const ComposeFolder = "/etc/sr/compose"

func ListComposes() ([]string, error) {
	files, err := ioutil.ReadDir(ComposeFolder)
	if err != nil {
		return nil, err
	}

	composes := make([]string, 0)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yaml") && !f.IsDir() {
			composes = append(composes, strings.TrimSuffix(f.Name(), ".yaml"))
		}
	}
	return composes, nil
}

func GetComposePath(name string) string {
	return ComposeFolder + "/" + name + ".yaml"
}

func IsComposeExist(name string) bool {
	if _, err := os.Stat(GetComposePath(name)); os.IsNotExist(err) {
		return false
	}
	return true
}

func RemoveCompose(name string) error {
	err := os.Remove(GetComposePath(name))
	if err != nil {
		return err
	}
	return nil
}

func WriteCompose(name string, compose models.Compose) error {
	file, err := os.Create(GetComposePath(name))
	if err != nil {
		return err
	}

	defer func(open *os.File) {
		_ = open.Close()
	}(file)

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	marshal, err := yaml.Marshal(&compose)
	if err != nil {
		return err
	}

	log.Println(string(marshal))
	_, err = file.WriteString(string(marshal))
	if err != nil {
		return err
	}

	return nil
}

func StackToCompose(stack models.Stack) models.Compose {
	services := map[string]models.ComposeService{}
	for _, service := range stack.Services {
		services[service.Name] = ServiceToComposeService(service)
	}

	volumes := map[string]models.ComposeVolumeDeclaration{}
	for _, volume := range stack.Volumes {
		volumes[volume.Name] = models.ComposeVolumeDeclaration{}
	}

	networks := map[string]models.ComposeNetworkDeclaration{
		"traefik-public": {
			External: true,
		},
	}
	for _, network := range stack.Networks {
		networks[network.Name] = models.ComposeNetworkDeclaration{}
	}

	return models.Compose{
		Services: services,
		Volumes:  volumes,
		Networks: networks,
	}
}

func ServiceToComposeService(service models.Service) models.ComposeService {

	var ports []string
	for _, port := range service.Ports {
		input := strconv.FormatUint(uint64(port.InputPort), 10)
		output := strconv.FormatUint(uint64(port.OutputPort), 10)
		ports = append(ports, input+":"+output)
	}

	var networks []string
	for _, network := range service.Networks {
		networks = append(networks, network.Name)
	}

	var volumes []string
	for _, volume := range service.Volumes {
		volumes = append(volumes, volume.Name+":"+volume.MountPoint)
	}

	environment := map[string]string{}
	for _, env := range service.Envs {
		environment[env.Name] = env.Value
	}

	labels := map[string]string{}
	if service.Domain != "" {

		networks = AddOnce(networks, "traefik-public")

		tPrefix := "traefik.http.routers." + service.Domain

		labels = map[string]string{
			"traefik.enable":         "true",
			tPrefix + ".entrypoints": "websecure",
			tPrefix + ".rule":        "Host(`" + service.Domain + "`)",
			tPrefix + ".service":     service.Domain,
			"traefik.http.services." + service.Domain + ".loadbalancer.server.port": strconv.Itoa(service.HttpPort),
		}
	}

	return models.ComposeService{
		Image:       service.Image + ":" + service.ImageVersion,
		Ports:       ports,
		Networks:    networks,
		Volumes:     volumes,
		Environment: environment,
		Labels:      labels,
	}
}

func ReadComposeToStack(name string) (*models.Stack, error) {

	file, err := ioutil.ReadFile(GetComposePath(name))
	if err != nil {
		return nil, err
	}

	compose := models.Compose{}

	err = yaml.Unmarshal(file, &compose)
	if err != nil {
		return nil, err
	}

	stack, err := ComposeToStack(compose, name)
	if err != nil {
		return nil, err
	}

	return stack, nil
}

func ComposeToStack(compose models.Compose, name string) (*models.Stack, error) {
	networks, err := ComposeNetworksToNetworkDeclarations(compose.Networks)
	if err != nil {
		return nil, err
	}

	volumes, err := ComposeVolumesToVolumeDeclarations(compose.Volumes)
	if err != nil {
		return nil, err
	}

	services, err := ComposeServicesToServices(compose.Services)
	if err != nil {
		return nil, err
	}
	return &models.Stack{
		Networks: networks,
		Volumes:  volumes,
		Services: services,
		Name:     name,
	}, nil
}

func ComposeNetworksToNetworkDeclarations(networks map[string]models.ComposeNetworkDeclaration) ([]models.NetworkDeclaration, error) {
	var returnNetworks []models.NetworkDeclaration
	for name, _ := range networks {
		returnNetworks = append(returnNetworks, models.NetworkDeclaration{
			Name: name,
		})
	}

	return returnNetworks, nil
}

func ComposeVolumesToVolumeDeclarations(networks map[string]models.ComposeVolumeDeclaration) ([]models.VolumeDeclaration, error) {
	var returnVolumes []models.VolumeDeclaration
	for name, _ := range networks {
		returnVolumes = append(returnVolumes, models.VolumeDeclaration{
			Name: name,
		})
	}

	return returnVolumes, nil
}

func ComposeServicesToServices(services map[string]models.ComposeService) ([]models.Service, error) {
	var returnService []models.Service
	for name, service := range services {
		ports, err := StringsToPorts(service.Ports)

		if err != nil {
			return nil, err
		}

		envs, err := StringsToEnvs(service.Environment)
		if err != nil {
			return nil, err
		}

		vol, err := StringsToVolumes(service.Volumes)
		if err != nil {
			return nil, err
		}

		networks, err := StringsToNetworks(service.Networks)
		if err != nil {
			return nil, err
		}

		imageInfo := strings.Split(service.Image, ":")
		imageVersion := "latest"
		if len(imageInfo) > 1 {
			imageVersion = imageInfo[1]
		}

		httpPort := 0
		domain := ""

		for label, value := range service.Labels {
			if strings.HasSuffix(label, "loadbalancer.server.port") {
				httpPort, err = strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
			} else if strings.HasSuffix(label, ".service") {
				domain = value
			}
		}

		returnService = append(returnService, models.Service{
			Name:         name,
			Image:        imageInfo[0],
			ImageVersion: imageVersion,
			Ports:        ports,
			Envs:         envs,
			Volumes:      vol,
			Networks:     networks,
			Domain:       domain,
			HttpPort:     httpPort,
		})
	}

	return returnService, nil
}

func StringsToPorts(ports []string) ([]models.Port, error) {
	var returnPorts []models.Port
	for _, value := range ports {
		parts := strings.Split(value, ":")
		input, err := strconv.ParseUint(parts[0], 10, 32)
		output, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			return nil, err
		}
		returnPorts = append(returnPorts, models.Port{
			InputPort:  uint(input),
			OutputPort: uint(output),
		})
	}

	return returnPorts, nil
}

func StringsToEnvs(ports map[string]string) ([]models.EnvVar, error) {
	var returnEnvs []models.EnvVar
	for name, value := range ports {
		returnEnvs = append(returnEnvs, models.EnvVar{
			Name:  name,
			Value: value,
		})
	}

	return returnEnvs, nil
}

func StringsToVolumes(ports []string) ([]models.Volume, error) {
	var returnVolumes []models.Volume
	for _, value := range ports {
		parts := strings.Split(value, ":")

		returnVolumes = append(returnVolumes, models.Volume{
			Name:       parts[0],
			MountPoint: parts[0],
		})
	}

	return returnVolumes, nil
}

func StringsToNetworks(ports []string) ([]models.Network, error) {
	var returnNetworks []models.Network
	for _, value := range ports {
		returnNetworks = append(returnNetworks, models.Network{
			Name: value,
		})
	}

	return returnNetworks, nil
}
