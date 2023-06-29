package middleware

import (
	"brainwave/internal/global"
	"brainwave/pkg/consts/berr"
	"brainwave/pkg/util/helper"
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
)

func Handler(handleFunc func(c *gin.Context) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		funcName := runtime.FuncForPC(reflect.ValueOf(handleFunc).Pointer()).Name()
		res, err := handleFunc(c)
		er := berr.DecodeErr(err)
		if er != nil {
			global.LOG.Error(ctx, funcName, er)
			helper.SetResponse(c, er, nil)
			return
		}
		helper.SetResponse(c, nil, res)
	}
}
