package service

import (
	"brainwave/internal/dto/request"
	"brainwave/internal/dto/response"
	"github.com/gin-gonic/gin"
)

type IUser interface {
	Login(c *gin.Context, req request.LoginReq) (*response.LoginRes, error)
}
