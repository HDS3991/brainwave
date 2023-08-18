package user

import (
	"brainwave/internal/dto/request"
	"brainwave/internal/dto/response"
	"brainwave/internal/service"
	"github.com/gin-gonic/gin"
)

type Service struct{}

func NewUser() service.IUser {
	return &Service{}
}

func (s *Service) Login(c *gin.Context, req request.LoginReq) (*response.LoginRes, error) {
	return nil, nil
}
