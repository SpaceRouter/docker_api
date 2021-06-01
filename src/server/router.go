package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spacerouter/docker_api/controllers"
	_ "github.com/spacerouter/docker_api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	router.POST("/tea", controllers.GetTea)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
