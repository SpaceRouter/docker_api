package forms

import (
	"github.com/docker/docker/api/types"
	"github.com/spacerouter/docker_api/models"
)

type StackResponse struct {
	Message   string
	Ok        bool
	Stack     *models.Stack
	Developer *models.Developer
}

type BasicResponse struct {
	Message string
	Ok      bool
}

type StackListResponse struct {
	Message string
	Ok      bool
	Stacks  []string
}

type ContainersResponse struct {
	Message    string
	Ok         bool
	Containers []types.Container `swaggerignore:"true"`
}

type StackInfo struct {
	ID          uint
	Name        string
	Icon        string
	Description string
	Developer   *models.Developer
}
