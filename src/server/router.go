package server

import (
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/spacerouter/docker_api/controllers"
	_ "github.com/spacerouter/docker_api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(cli *client.Client) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	router.POST("/tea", controllers.GetTea)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//auth := sr_auth.CreateAuth(config.GetSecretKey(), config.GetConfig().GetString("security.auth_server"), nil)

	v1 := router.Group("v1")
	{
		dockerCtrl := controllers.DockerController{
			Client: cli,
		}
		//v1.Use(auth.SrAuthMiddlewareGin())
		v1.GET("containers", dockerCtrl.GetContainers)

	}

	return router
}
