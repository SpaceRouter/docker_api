package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func StartStack(stackName string) error {

	cmd := exec.Command("docker-compose", "-f", GetComposePath(stackName), "-p", stackName, "up", "-d")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	errStr := stderr.String()
	if errStr != "" {
		return fmt.Errorf(errStr)
	}

	return nil
}

func StopStack(stackName string) error {

	cmd := exec.Command("docker-w", "-f", GetComposePath(stackName), "down")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	errStr := stderr.String()
	if errStr != "" {
		return fmt.Errorf(errStr)
	}

	return nil
}

// GetComposeContainerIds get all child containers ids of all composes
func GetComposeContainerIds() ([]string, error) {

	cmd := exec.Command("docker", "ps", "--filter", "label=com.docker.compose.project", "-q")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	errStr := stderr.String()
	if errStr != "" {
		return nil, fmt.Errorf(errStr)
	}

	ids := strings.Split(stdout.String(), "\n")
	return ids, nil
}
