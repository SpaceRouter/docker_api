package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTea(c *gin.Context) {
	c.JSON(http.StatusTeapot, gin.H{
		"Message": "I'm a teapot",
		"Ok":      true,
		"Token":   "UWANNADRINKTEA",
	})
	return
}
