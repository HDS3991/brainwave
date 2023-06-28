package server

import (
	"brainwave/internal/global"
	"brainwave/internal/init/log"
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

	address := fmt.Sprintf("%s:%s", global.Conf.System.Host, global.Conf.System.HttpPort)
	s := endless.NewServer(address, nil)
	s.ReadHeaderTimeout = readHeaderTimeoutSecond * time.Second
	s.WriteTimeout = writeTimeoutSecond * time.Second
	s.MaxHeaderBytes = maxHeaderBytes

	if global.Conf.System.SSL == sslDisable {
		global.LOG.Infof("server run success on %s:%s with http", global.Conf.System.Host, global.Conf.System.HttpPort)
		if err := s.ListenAndServe(); err != nil {
			global.LOG.Error(err)
			panic(err)
		}

	}

}
