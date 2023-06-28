package logger

import (
	"errors"
	"io"
	"os"
	"path"
)

var (
	BufferSize         = 0x100000
	DefaultFileMode    = os.FileMode(0644)
	DefaultFileFlag    = os.O_RDWR | os.O_CREATE | os.O_APPEND
	ErrInvalidArgument = errors.New("error argument invalid")
	QueueSize          = 1024
	ErrClosed          = errors.New("error write on close")
)

type Config struct {
	TimeTagFormat      string
	LogPath            string
	FileName           string
	LogSuffix          string
	MaxRemain          int
	RollingTimePattern string
}

type RollingWriter interface {
	io.Writer
	Close() error
}

type Manager interface {
	Fire() chan string
	Close()
}

func FilePath(c *Config) (filepath string) {
	filepath = path.Join(c.LogPath, c.FileName) + c.LogSuffix
	return
}
