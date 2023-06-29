package base

import (
	"brainwave/internal/app/handler/v1"
	"brainwave/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (s *Router) InitBaseRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("auth")
	baseApi := v1.ApiGroupApp.BaseAPi
	{
		baseRouter.POST("/login", middleware.Handler(baseApi.Login))
	}
}
