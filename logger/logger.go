package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var sugarLogger *zap.SugaredLogger

// InitLogger initialize the zap logger
func InitLogger() {
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, getLogWriter(), zapcore.InfoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

// GetLogger return this logger instance
func GetLogger() *zap.SugaredLogger {
	return sugarLogger
}

// getEncoder returns the log encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter initialize and return the file logger instance
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
