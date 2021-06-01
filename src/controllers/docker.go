package controllers

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/spacerouter/docker_api/forms"
	"github.com/spacerouter/docker_api/models"
	"net/http"
)

type DockerController struct {
	Client *client.Client
}

func (dc *DockerController) GetContainers(c *gin.Context) {
	list, err := dc.Client.ContainerList(c, types.ContainerListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, forms.ContainersResponse{
			Ok:         false,
			Message:    err.Error(),
			Containers: nil,
		})
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
		c.JSON(http.StatusBadRequest, forms.ContainersResponse{
			Ok:         false,
			Message:    err.Error(),
			Containers: nil,
		})
	}

	// TODO Create stack

	list, err := dc.Client.ContainerList(c, types.ContainerListOptions{})
	if err != nil {
		c.JSON(http.StatusOK, forms.ContainersResponse{
			Ok:         false,
			Message:    err.Error(),
			Containers: nil,
		})
	}

	c.JSON(http.StatusOK, forms.ContainersResponse{
		Ok:         true,
		Message:    "",
		Containers: list,
	})
}
