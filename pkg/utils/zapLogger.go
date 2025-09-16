package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitSugaredLogger() error {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),     //or NewConsoleEncoder           
		zapcore.AddSync(logFile),                             
		zap.NewAtomicLevelAt(zapcore.InfoLevel),             
	)

	// ساخت logger
	logger := zap.New(core)
	Logger = logger.Sugar()

	return nil
}