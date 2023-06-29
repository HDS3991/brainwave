package v1

import (
	"brainwave/internal/app/handler/v1/base"
	baseSrv "brainwave/internal/app/service/base"
)

type ApiGroup struct {
	BaseAPi base.Base
}

var ApiGroupApp = new(ApiGroup)

var (
	baseService = baseSrv.NewIService()
)
