package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func LoggerInit() {
	l, err := zap.NewProduction()
	if err != nil {
		return
	}
	Logger = l
}
