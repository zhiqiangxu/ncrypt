package logger

import (
	"fmt"

	l "github.com/zhiqiangxu/util/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New is ctor for zap.Logger
func New() *zap.Logger {

	lvl := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	var encoderConfig zapcore.EncoderConfig
	encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	zconf := zap.Config{
		DisableCaller:     true,
		DisableStacktrace: true,
		Level:             lvl,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}

	logger, err := l.New(zconf)
	if err != nil {
		panic(fmt.Sprintf("Build:%v", err))
	}

	return logger
}
