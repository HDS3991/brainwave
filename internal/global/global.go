package global

import (
	"brainwave/internal/conf/server"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	CONF  server.Conf
	VIPER *viper.Viper
	LOG   *logrus.Logger
	DB    *gorm.DB
	VALID *validator.Validate
)
