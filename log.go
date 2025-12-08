package log

import (
	"github.com/go-think/log/record"
)

var logger *Logger

func init() {
	logger = NewLogger("develop", record.DEBUG)
}

// Debug Adds a log record at the DEBUG level.
func Debug(format string, v ...interface{}) {
	logger.Debug(format, v...)
}

// Info Adds a log record at the INFO level.
func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

// Notice Adds a log record at the NOTICE level.
func Notice(format string, v ...interface{}) {
	logger.Notice(format, v...)
}

// Warn Adds a log record at the WARNING level.
func Warn(format string, v ...interface{}) {
	logger.Warn(format, v...)
}

// Error Adds a log record at the ERROR level.
func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}

// Crit Adds a log record at the CRITICAL level.
func Crit(format string, v ...interface{}) {
	logger.Crit(format, v...)
}

// Alert Adds a log record at the ALERT level.
func Alert(format string, v ...interface{}) {
	logger.Alert(format, v...)
}

// Emerg Adds a log record at the EMERGENCY level.
func Emerg(format string, v ...interface{}) {
	logger.Emerg(format, v...)
}

// GetLogger Get the default Logger
func GetLogger() *Logger {
	return logger
}

// SetLogger Set the default Logger
func SetLogger(l *Logger) {
	logger = l
}
