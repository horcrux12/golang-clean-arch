package helper

import (
	"github.com/horcrux12/clean-rest-api-template/constanta"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type logger struct {
	LoggerDebugLevel *zap.Logger
}

var Logging logger

func ConfigZap(level zapcore.Level, output []string) (logger *zap.Logger) {
	cfg := zap.Config{
		Encoding:    "json",
		OutputPaths: output,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:     "timestamp",
			EncodeTime:  zapcore.RFC3339NanoTimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}
	cfg.Level = zap.NewAtomicLevelAt(level)

	logger, err := cfg.Build()
	if err != nil {
		os.Exit(9)
	}
	return
}

func SetLogger(logFileName []string, argument string) {

	switch argument {
	case constanta.DevelopmentArgument:
		Logging.LoggerDebugLevel = ConfigZap(zap.DebugLevel, logFileName)
		break
	default:
		Logging.LoggerDebugLevel = ConfigZap(zap.InfoLevel, logFileName)
		break
	}
}

func LogInfo(data []zap.Field) {
	Logging.LoggerDebugLevel.Info("", data...)
}

func LogError(data []zap.Field) {
	Logging.LoggerDebugLevel.Error("", data...)
}
