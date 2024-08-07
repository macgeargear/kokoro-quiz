package logger

import (
	"github.com/macgeargear/kokoro-quiz/internal/config"
	"go.uber.org/zap"
)

func New(conf *config.Config) *zap.Logger {
	var logger *zap.Logger

	if conf.App.IsDevelopment() {
		logger = zap.Must(zap.NewDevelopment())
	} else {
		logger = zap.Must(zap.NewProduction())
	}

	return logger
}
