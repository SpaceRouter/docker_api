package controllers

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/spacerouter/docker_api/forms"
	"github.com/spacerouter/docker_api/models"
	"github.com/spacerouter/docker_api/utils"
	"log"
	"net/http"
)

type DockerController struct {
	Client *client.Client
}

// GetContainers godoc
// @Summary Get all containers
// @Description Get all containers
// @ID GetContainers
// @Produce  json
// @Success 200 {object} forms.ContainersResponse
// @Failure 500 {object} forms.ContainersResponse
// @Router /v1/containers [get]
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

// CreateStack godoc
// @Summary Create or modify stack
// @Description Create or modify stack
// @ID CreateStack
// @Produce  json
// @Accept  json
// @Param Stack body models.Stack false "Stack"
// @Success 200 {object} forms.BasicResponse
// @Failure 500,400,401 {object} forms.BasicResponse
// @Router /v1/stack/ [post]
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
	log.Println()
	err = utils.WriteCompose(stack.Name, compose)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, forms.BasicResponse{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, forms.BasicResponse{
		Ok:      true,
		Message: "",
	})
}

// GetStack godoc
// @Summary Get stack info
// @Description Get stack info
// @ID GetStack
// @Produce  json
// @Success 200 {object} forms.StackResponse
// @Failure 500 {object} forms.StackResponse
// @Router /v1/stack/{name} [get]
func (dc *DockerController) GetStack(c *gin.Context) {
	stack, err := utils.ReadComposeToStack(c.Param("name"))
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

// GetStackList godoc
// @Summary Get stack list
// @Description Get stack list
// @ID GetStackList
// @Produce  json
// @Success 200 {object} forms.StackListResponse
// @Failure 500 {object} forms.StackListResponse
// @Router /v1/stacks [get]
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

// StartStack godoc
// @Summary Start stack
// @Description Start stack
// @ID StartStack
// @Produce  json
// @Success 200 {object} forms.BasicResponse
// @Failure 500,404 {object} forms.BasicResponse
// @Router /v1/stack/{name}/start [get]
func (dc *DockerController) StartStack(c *gin.Context) {
	name := c.Param("name")
	exist := utils.IsComposeExist(name)
	if !exist {
		c.AbortWithStatusJSON(http.StatusNotFound, forms.BasicResponse{
			Ok:      false,
			Message: "Stack not found",
		})
		return
	}

	err := utils.StartStack(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, forms.BasicResponse{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, forms.BasicResponse{
		Ok:      true,
		Message: "",
	})
}

// StopStack godoc
// @Summary Stop stack
// @Description Stop stack
// @ID StopStack
// @Produce  json
// @Success 200 {object} forms.BasicResponse
// @Failure 500,404 {object} forms.BasicResponse
// @Router /v1/stack/{name}/stop [get]
func (dc *DockerController) StopStack(c *gin.Context) {
	name := c.Param("name")
	exist := utils.IsComposeExist(name)
	if !exist {
		c.AbortWithStatusJSON(http.StatusNotFound, forms.BasicResponse{
			Ok:      false,
			Message: "Stack not found",
		})
		return
	}

	err := utils.StopStack(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, forms.BasicResponse{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, forms.BasicResponse{
		Ok:      true,
		Message: "",
	})
}

// RemoveStack godoc
// @Summary Remove stack
// @Description Remove stack
// @ID RemoveStack
// @Produce  json
// @Success 200 {object} forms.BasicResponse
// @Failure 500,404 {object} forms.BasicResponse
// @Router /v1/stack/{name} [delete]
func (dc *DockerController) RemoveStack(c *gin.Context) {
	name := c.Param("name")
	exist := utils.IsComposeExist(name)
	if !exist {
		c.AbortWithStatusJSON(http.StatusNotFound, forms.BasicResponse{
			Ok:      false,
			Message: "Stack not found",
		})
		return
	}

	err := utils.StopStack(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, forms.BasicResponse{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	err = utils.RemoveCompose(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, forms.BasicResponse{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, forms.BasicResponse{
		Ok:      true,
		Message: "",
	})
}

// GetActiveStacks godoc
// @Summary Get running stacks
// @Description Get running stacks
// @ID GetActiveStacks
// @Produce  json
// @Success 200 {object} forms.ActiveStacksResponse
// @Failure 500,404 {object} forms.ActiveStacksResponse
// @Router /v1/active_stacks [get]
func (dc *DockerController) GetActiveStacks(c *gin.Context) {
	ids, err := utils.GetComposeContainerIds()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, forms.ActiveStacksResponse{
			Ok:      false,
			Message: err.Error(),
		})
		return
	}

	var composeList []string

	for _, id := range ids {
		if id == "" {
			continue
		}

		info, err := dc.Client.ContainerInspect(c, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, forms.ActiveStacksResponse{
				Ok:      false,
				Message: err.Error(),
			})
			return
		}

		composeList = utils.AddOnce(composeList, info.Config.Labels["com.docker.compose.project"])
	}
	c.JSON(http.StatusOK, forms.ActiveStacksResponse{
		Ok:      true,
		Message: "",
		Stacks:  composeList,
	})
}
