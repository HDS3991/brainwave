package router

import (
	"brainwave/internal/middleware"
	"brainwave/internal/router"
	"brainwave/pkg/i18n"
	"github.com/gin-contrib/gzip"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	routers := gin.Default()
	routers.Use(middleware.OperationLog())
	routers.Use(gzip.Gzip(gzip.DefaultCompression))
	routers.Use(i18n.GinI18nLocalize())
	publicGroup := routers.Group("")
	{
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	routerGroup := router.Group

	privateGroup := routers.Group("api/v1")
	privateGroup.Use(middleware.JwtAuth()).Use(middleware.SessionAuth())
	{
		routerGroup.InitBaseRouter(privateGroup)
	}

	return routers
}
