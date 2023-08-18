package controller

import (
	"brainwave/internal/dto/request"
	"brainwave/internal/global"
	"brainwave/internal/service"
	"brainwave/pkg/berr"
	"github.com/gin-gonic/gin"
)

type IUser interface {
	Login(c *gin.Context) (any, error)
}

func NewUser() IUser {
	return &User{}
}

type User struct{}

// Login
// @Tags User
// @Summary User login
// @Description 用户登录
// @Accept json
// @Param request body request.LoginReq true "request"
// @Success 200 {object} response.LoginRes
// @Router /api/user/login [POST]
func (u *User) Login(c *gin.Context) (any, error) {
	var req request.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, berr.NewErr(berr.ErrorInvalidArgument).Wrap(err)
	}
	if err := global.VALID.Struct(req); err != nil {
		return nil, berr.NewErr(berr.ErrorInvalidArgument).Wrap(err)
	}
	return service.Entry.User.Login(c, req)
}
