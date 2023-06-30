package helper

import (
	"brainwave/pkg/consts/berr"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func SetResponse(c *gin.Context, err berr.ErrI, data any) {
	if err != nil {
		msg := err.Msg()
		if gin.Mode() != gin.ReleaseMode {
			msg = err.Error()
		}
		c.JSON(http.StatusOK, Response{
			Code: err.Code(),
			Msg:  msg,
			Data: struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "",
		Data: data,
	})
}
