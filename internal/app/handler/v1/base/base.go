package base

import (
	"brainwave/internal/app/dto/request"
	"brainwave/pkg/consts/berr"
	"github.com/gin-gonic/gin"
)

type Base struct{}

// Login
// @Tags Auth
// @Summary User login
// @Description 用户登录
// @Accept json
// @Param request body request.LoginReq true "request"
// @Success 200 {object} response.LoginRes
// @Router api/v1/auth/login [post]
func (s *Base) Login(c *gin.Context) (any, error) {
	var req request.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, berr.NewErr(berr.ErrorInvalidArgument).Wrap(err)
	}
	return "ok", nil
}
