package router

import (
	"brainwave/internal/middleware"
	"brainwave/internal/router"
	"brainwave/pkg/i18n"
	"github.com/gin-contrib/gzip"

	"github.com/gin-gonic/gin"
)

type Routers struct {
	User router.IUser
}

var routers = Routers{
	User: router.NewUser(),
}

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Log())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(i18n.GinI18nLocalize())
	publicGroup := r.Group("")
	{
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	privateGroup := r.Group("/api")
	privateGroup.Use(middleware.JwtAuth())
	privateGroup.Use(middleware.SessionAuth())
	{
		routers.User.InitUserRouter(privateGroup)
	}

	return r
}
