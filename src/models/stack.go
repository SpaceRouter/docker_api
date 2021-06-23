package models

type Stack struct {
	Name     string
	Services []Service
	Networks []NetworkDeclaration
	Volumes  []VolumeDeclaration
}

type Service struct {
	Name         string
	Image        string
	ImageVersion string
	Domain       string
	HttpPort     int
	Ports        []Port
	Envs         []EnvVar
	Volumes      []Volume
	Networks     []Network
}

type Port struct {
	InputPort  uint
	OutputPort uint
}

type VolumeDeclaration struct {
	Name string
}

type Volume struct {
	Name       string
	MountPoint string
}

type EnvVar struct {
	Name  string
	Value string
}

type NetworkDeclaration struct {
	Name string
}

type Network struct {
	Name string
}
