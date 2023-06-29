package router

import (
	"brainwave/internal/middleware"
	"brainwave/internal/router"
	"github.com/gin-contrib/gzip"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.OperationLog())
	Router.Use(gzip.Gzip(gzip.DefaultCompression))

	publicGroup := Router.Group("")
	{
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	routers := router.Group

	privateGroup := Router.Group("api/v1")
	privateGroup.Use(middleware.JwtAuth()).Use(middleware.SessionAuth())
	{
		routers.InitBaseRouter(privateGroup)
	}

	return Router
}
