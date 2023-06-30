package server

import (
	"brainwave/internal/global"
	"brainwave/internal/init/db"
	"brainwave/internal/init/log"
	"brainwave/internal/init/router"
	"brainwave/internal/init/validator"
	"brainwave/internal/init/viper"
	"fmt"
	"github.com/fvbock/endless"
	"time"
)

const (
	readHeaderTimeoutSecond = 20
	writeTimeoutSecond      = 60
	maxHeaderBytes          = 1 << 20
	sslDisable              = "disable"
)

func Start() {
	viper.Init()
	log.Init()
	db.Init()
	validator.Init()

	handler := router.Init()

	address := fmt.Sprintf("%s:%s", global.CONF.System.Host, global.CONF.System.HttpPort)
	s := endless.NewServer(address, handler)
	s.ReadHeaderTimeout = readHeaderTimeoutSecond * time.Second
	s.WriteTimeout = writeTimeoutSecond * time.Second
	s.MaxHeaderBytes = maxHeaderBytes

	if global.CONF.System.SSL == sslDisable {
		global.LOG.Infof("server run success on %s:%s with http", global.CONF.System.Host, global.CONF.System.HttpPort)
		if err := s.ListenAndServe(); err != nil {
			global.LOG.Error(err)
			panic(err)
		}
	}

}
