package viper

import (
	"brainwave/internal/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

func Init() {
	v := viper.NewWithOptions()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.AddConfigPath("internal/etc/server/")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file not found")
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
	global.Viper = v
	if err := v.Unmarshal(&global.Conf); err != nil {
		panic(fmt.Errorf("config file init: %w", err))
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.Conf); err != nil {
			panic(fmt.Errorf("config file changed: %w", err))
		}
		time.Sleep(1 * time.Second)
	})
	v.WatchConfig()
	time.Sleep(1 * time.Second)
}
