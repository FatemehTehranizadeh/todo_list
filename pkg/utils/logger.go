package utils

import (
	"log"
	"os"
	// "go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	// zapLogger     *zap.SugaredLogger
)

func InitLogger() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		ErrorLogger.Println("Error while opening the log file: ", err)
	}
	defer file.Close()
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// func InitSugaredLogger() error {
// 	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	encoderConfig := zap.NewProductionEncoderConfig()
// 	encoderConfig.TimeKey = "time"
// 	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

// 	core := zapcore.NewCore(
// 		zapcore.NewJSONEncoder(encoderConfig),
// 		zapcore.AddSync(logFile),
// 		zap.NewAtomicLevelAt(zapcore.InfoLevel),
// 	)

// 	logger := zap.New(core)
// 	zapLogger = logger.Sugar()

// 	return nil
// }
