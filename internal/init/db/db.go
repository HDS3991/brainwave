package db

import (
	"brainwave/internal/global"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Init() {
	if _, err := os.Stat(global.Conf.System.DB.Path); err != nil {
		if err := os.MkdirAll(global.Conf.System.DB.Path, os.ModePerm); err != nil {
			panic(fmt.Errorf("init db dir falied, err: %v", err))
		}
	}
	fullPath := global.Conf.System.DB.Path + "/" + global.Conf.System.DB.FileName
	if _, err := os.Stat(fullPath); err != nil {
		if _, err := os.Create(fullPath); err != nil {
			panic(fmt.Errorf("init db file falied, err: %v", err))
		}
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})
	if err != nil {
		panic(fmt.Errorf("init db failed, err: %v", err))
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("init db failed, err: %v", err))
	}
	sqlDb.SetMaxIdleConns(global.Conf.System.DB.MaxIdleConns)
	sqlDb.SetMaxOpenConns(global.Conf.System.DB.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(global.Conf.System.DB.MaxLifeTimeSecond) * time.Second)

	global.DB = db
	global.LOG.Info(fmt.Sprintf("init db success, db file: %s", fullPath))
}
