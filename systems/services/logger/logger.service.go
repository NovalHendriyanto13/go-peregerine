package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	Log *zap.Logger
}

func Build() (*Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logger{
		Log: log,
	}, nil
}