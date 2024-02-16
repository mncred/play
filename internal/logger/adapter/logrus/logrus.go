// package logrus provides Wails logger.Logger adapter for logrus.Logger
package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

type Logger struct {
	logger *logrus.Entry
}

func Provide(logger *logrus.Entry) logger.Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Print(message string) {
	l.logger.Print(message)
}

func (l *Logger) Trace(message string) {
	l.logger.Trace(message)
}

func (l *Logger) Debug(message string) {
	l.logger.Debug(message)
}

func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

func (l *Logger) Warning(message string) {
	l.logger.Warn(message)
}

func (l *Logger) Error(message string) {
	l.logger.Error(message)
}

func (l *Logger) Fatal(message string) {
	l.logger.Fatal(message)
}
