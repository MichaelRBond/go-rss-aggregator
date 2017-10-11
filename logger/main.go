package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/michaelrbond/go-rss-aggregator/configuration"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

type zapLogger func(msg string, fields ...zapcore.Field)

func log(f zapLogger, msg string, level string) {
	f(
		msg,
		zap.String("level", level),
		zap.String("date", time.Now().Format("2006-01-02 15:04:05.000")),
	)
}

// Info logs an informational message
func Info(msg string) {
	GetLogger()
	log(logger.Info, msg, "info")
}

// Debug logs a debug message
func Debug(msg string) {
	GetLogger()
	log(logger.Debug, msg, "debug")
}

// Error logs an error message
func Error(msg string) {
	GetLogger()
	log(logger.Error, msg, "error")
}

// Panic logs an error message and then terminates the program
func Panic(msg string) {
	GetLogger()
	log(logger.Panic, msg, "panic")
}

// GetLogger gets. a. logger.
func GetLogger() {
	if logger != nil {
		return
	}
	config := configuration.GetConfig()
	if l, err := config.Logger.Build(); err != nil {
		fmt.Fprintf(os.Stderr, "Error getting logger: %s\n", err.Error())
		os.Exit(1)
	} else {
		logger = l
	}
}
