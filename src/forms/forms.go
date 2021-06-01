package forms

import "github.com/spacerouter/docker_api/models"

type StackResponse struct {
	Message   string
	Ok        bool
	Stack     *models.Stack
	Developer *models.Developer
}

type StackSearchResponse struct {
	Message string
	Ok      bool
	Stacks  []StackInfo
}

type StackInfo struct {
	ID          uint
	Name        string
	Icon        string
	Description string
	Developer   *models.Developer
}
