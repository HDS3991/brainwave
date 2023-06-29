package base

import (
	"brainwave/internal/app/dto/request"
	"brainwave/internal/app/dto/response"
	"github.com/gin-gonic/gin"
)

type Service struct{}

type IService interface {
	Login(c *gin.Context, req request.LoginReq) (*response.LoginRes, error)
}

func NewIService() IService {
	return &Service{}
}

func (s *Service) Login(c *gin.Context, req request.LoginReq) (*response.LoginRes, error) {
	return nil, nil
}
