package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func StartStack(stackName string) error {

	cmd := exec.Command("docker-compose", "up", "-f", GetComposePath(stackName), "-d")

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

	cmd := exec.Command("docker-compose", "down", "-f", GetComposePath(stackName))

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
