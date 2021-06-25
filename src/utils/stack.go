package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func StartStack(stackName string) error {
	dockerComposePath, err := exec.LookPath("docker-compose")
	if err != nil {
		return err
	}

	cmd := exec.Command(dockerComposePath, "-f", GetComposePath(stackName), "-p", stackName, "up", "-d")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		errStr := stderr.String()
		return fmt.Errorf("%s %s", err, errStr)
	}

	return nil
}

func StopStack(stackName string) error {
	dockerComposePath, err := exec.LookPath("docker-compose")
	if err != nil {
		return err
	}

	cmd := exec.Command(dockerComposePath, "-p", stackName, "-f", GetComposePath(stackName), "down")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		errStr := stderr.String()
		return fmt.Errorf("%s %s", err, errStr)
	}

	return nil
}

// GetComposeContainerIds get all child containers ids of all composes
func GetComposeContainerIds() ([]string, error) {
	dockerPath, err := exec.LookPath("docker")
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(dockerPath, "ps", "--filter", "label=com.docker.compose.project", "-q")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err = cmd.Run()
	if err != nil {
		errStr := stderr.String()
		return nil, fmt.Errorf("%s %s", err, errStr)
	}

	ids := strings.Split(stdout.String(), "\n")
	return ids, nil
}
