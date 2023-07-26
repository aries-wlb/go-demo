package logger

import (
	"golang.org/x/exp/slog"
)

func Info(msg string) {
	slog.Info(msg)
}

func Error(msg string) {
	slog.Error(msg)
}

func Warn(msg string) {
	slog.Warn(msg)
}

func Debug(msg string) {
	slog.Debug(msg)
}
