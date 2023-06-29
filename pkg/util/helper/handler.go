package helper

import (
	"brainwave/pkg/consts/berr"
	"fmt"
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
		fmt.Println(err.Code(), err.Error())
		c.JSON(http.StatusOK, Response{
			Code: err.Code(),
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "",
		Data: data,
	})
}
