package global

import (
	"brainwave/internal/conf/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Conf  server.Conf
	Viper *viper.Viper
	LOG   *logrus.Logger
	DB    *gorm.DB
)
