package utils

import (
	"bytes"
	"fmt"
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
