package log

import (
	"brainwave/internal/conf/server"
	"brainwave/internal/global"
	"brainwave/pkg/consts/system"
	"brainwave/pkg/logger"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	RollingTimePattern = "0 0  * * *"
)

func Init() {
	l := logrus.New()
	setOutput(l, global.CONF.Log)
	global.LOG = l
	global.LOG.Info("init logger successfully")
}

func setOutput(log *logrus.Logger, config server.LogConfig) {
	writer, err := logger.NewWriterFromConfig(&logger.Config{
		LogPath:            global.CONF.System.LogPath,
		FileName:           config.LogName,
		TimeTagFormat:      system.DateFormat,
		MaxRemain:          config.MaxBackup,
		RollingTimePattern: RollingTimePattern,
		LogSuffix:          config.LogSuffix,
	})
	if err != nil {
		panic(err)
	}
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		panic(err)
	}
	fileAndStdoutWriter := io.MultiWriter(writer, os.Stdout)

	log.SetOutput(fileAndStdoutWriter)
	log.SetLevel(level)
	log.SetFormatter(new(MineFormatter))
}

type MineFormatter struct{}

func (s *MineFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	detailInfo := ""
	if entry.Caller != nil {
		fc := entry.Caller.Function
		file := entry.Caller.File
		detailInfo = fmt.Sprintf("(%v-%s: %d)", file, fc, entry.Caller.Line)
	}
	msg := fmt.Sprintf("[%s] [%s] {%s} %s {%v} \n",
		entry.Time.Format(system.DateTimeFormat),
		entry.Level.String(),
		entry.Message,
		detailInfo,
		entry.Data)
	return []byte(msg), nil
}
