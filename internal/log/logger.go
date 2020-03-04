package log

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Logger struct {
	*logrus.Logger
}

var (
	instance *Logger
	once     sync.Once
)

func newLogger() *Logger {
	log := Logger{logrus.New()}
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		ForceColors:     true,
	}
	return &log
}

func GetLogger() *Logger {
	once.Do(func() {
		instance = newLogger()
	})
	return instance
}
