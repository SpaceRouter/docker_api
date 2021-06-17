package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spacerouter/docker_api/models"
	"os/exec"
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

	cmd := exec.Command("docker-compose", "-f", GetComposePath(stackName), "down")

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

func GetActiveStacks() ([]models.ComposeOutput, error) {

	cmd := exec.Command("docker", "compose", "ls", "--format", "json")

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

	var stacks []models.ComposeOutput
	err = json.Unmarshal(stdout.Bytes(), &stacks)
	if err != nil {
		return nil, err
	}

	return stacks, nil
}
