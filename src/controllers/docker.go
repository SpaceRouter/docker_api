package controllers

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/spacerouter/docker_api/forms"
	"github.com/spacerouter/docker_api/models"
	"github.com/spacerouter/docker_api/utils"
	"net/http"
)

type DockerController struct {
	Client *client.Client
}

func (dc *DockerController) GetContainers(c *gin.Context) {
	list, err := dc.Client.ContainerList(c, types.ContainerListOptions{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, forms.ContainersResponse{
			Ok:         false,
			Message:    err.Error(),
			Containers: nil,
		})
		return
	}

	c.JSON(http.StatusOK, forms.ContainersResponse{
		Ok:         true,
		Message:    "",
		Containers: list,
	})
}

func (dc *DockerController) CreateStack(c *gin.Context) {
	var stack models.Stack

	err := c.ShouldBindJSON(&stack)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, forms.BasicResponse{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	compose := utils.StackToCompose(stack)
	err = utils.WriteCompose(stack.Name, compose)
	if err != nil {
		c.JSON(http.StatusInternalServerError, forms.BasicResponse{
			Ok:      false,
			Message: err.Error(),
		})
	}
}

func (dc *DockerController) GetStack(c *gin.Context) {
	stack, err := utils.ImportCompose(c.Param("name"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, forms.StackResponse{
			Ok:      false,
			Message: err.Error(),
			Stack:   stack,
		})
		return
	}

	c.JSON(http.StatusOK, forms.StackResponse{
		Ok:      true,
		Message: "",
		Stack:   stack,
	})
}

func (dc *DockerController) GetStackList(c *gin.Context) {
	stacks, err := utils.ListComposes()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, forms.StackListResponse{
			Ok:      false,
			Message: err.Error(),
			Stacks:  nil,
		})
		return
	}

	c.JSON(http.StatusOK, forms.StackListResponse{
		Ok:      true,
		Message: "",
		Stacks:  stacks,
	})
}
