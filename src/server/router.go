package server

import (
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/spacerouter/docker_api/controllers"
	_ "github.com/spacerouter/docker_api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(cli *client.Client) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	health := new(controllers.HealthController)

	main := router.Group("docker")
	{

		main.GET("/health", health.Status)

		main.POST("/tea", controllers.GetTea)

		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		//auth := sr_auth.CreateAuth(config.GetSecretKey(), config.GetConfig().GetString("security.auth_server"), nil)

		v1 := main.Group("v1")
		{
			dockerCtrl := controllers.DockerController{
				Client: cli,
			}
			//v1.Use(auth.SrAuthMiddlewareGin())
			v1.GET("containers", dockerCtrl.GetContainers)
			v1.GET("stacks", dockerCtrl.GetStackList)
			v1.GET("active_stacks", dockerCtrl.GetActiveStacks)

			stack := v1.Group("stack")
			{
				stack.GET(":name", dockerCtrl.GetStack)
				stack.POST("", dockerCtrl.CreateStack)
				stack.DELETE(":name", dockerCtrl.RemoveStack)
				stack.GET(":name/start", dockerCtrl.StartStack)
				stack.GET(":name/stop", dockerCtrl.StopStack)
			}
		}
	}

	return router
}
