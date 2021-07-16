package init

import (
	"github.com/charlloss/go-rest-api/api"
	"github.com/charlloss/go-rest-api/routers"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Routers() *gin.Engine {
	router := gin.Default()

	router.Static("/public", "./public")
	//router.NoRoute(response.NotFound)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Teapot
	router.GET("teapot", func(c *gin.Context) {
		c.JSON(http.StatusTeapot, gin.H{
			"message": "I'm a teapot",
			"story": "This code was defined in 1998 " +
				"as one of the traditional IETF April Fools' jokes," +
				" in RFC 2324, Hyper Text Coffee Pot Control Protocol," +
				" and is not expected to be implemented by actual HTTP servers." +
				" However, known implementations do exist.",
		})
	})

	router.GET("/", api.GetApiList)

	publicRouterV1 := router.Group("v1")
	{
		routers.InitBaseRouter(publicRouterV1)
	}

	return router
}
