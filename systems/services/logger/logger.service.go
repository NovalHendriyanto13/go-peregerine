package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Logger struct {
	Log *zap.Logger
}

func Build() (*Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	appCore := zapcore.NewCore(
		fileEncoder,
		getFileWriter("storage/logs/app.log"),
		zap.InfoLevel,
	)

	consoleCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)

	core := zapcore.NewTee(appCore, consoleCore)
	log := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{
		Log: log,
	}, nil
}

func getFileWriter(filename string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename: filename,
		MaxSize: 20,
		MaxBackups: 5,
		MaxAge: 30,
		Compress: true,
	})
}