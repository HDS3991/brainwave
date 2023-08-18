package router

import (
	"brainwave/internal/controller"
	"brainwave/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewUser() IUser {
	return &User{}
}

type IUser interface {
	InitUserRouter(router *gin.RouterGroup)
}

type User struct{}

func (s *User) InitUserRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("/user")
	baseApi := controller.Entry.User
	{
		baseRouter.POST("/login", middleware.Handler(baseApi.Login))
	}
}
