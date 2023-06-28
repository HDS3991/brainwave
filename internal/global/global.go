package global

import (
	"brainwave/internal/conf/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Conf  server.Conf
	Viper *viper.Viper
	LOG   *logrus.Logger
)
