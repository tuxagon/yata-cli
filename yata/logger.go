package yata

import (
	"fmt"
	"os"
	"path/filepath"
)

const logFilename = ".yatalog"
const (
	LogVerbose = iota
	LogInfo
	LogWarning
	LogError
	LogNone
)

// LogLevel represents the lowest level of
// logging the user wants displayed during
// the execution of the app
var LogLevel = LogInfo

// Logger provides the contract for logging
// verbose, info, warning, and error messages
// using a pre-defined heirarchy
type Logger interface {
	Verbose(string)
	Info(string)
	Warning(string)
	Error(string)
	Write(int, string)
}

var logger Logger

// FileLogger TODO docs
type FileLogger struct {
	stdout bool
	file   *os.File
	lvl    int
}

// GetLogger gets the configured logger or if a logger is
// not set, will create a new FileLogger
func GetLogger() Logger {
	if logger != nil {
		return logger
	}

	path := filepath.Join(GetDirectory().MetadataDir(), "logs")
	HandleError(mkdir(path)(), false)
	filename := filepath.Join(path, logFilename)

	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	logger = &FileLogger{
		stdout: true,
		file:   fp,
		lvl:    LogLevel,
	}

	if err != nil {
		logger.Warning(err.Error())
	}

	return logger
}

// Verbose TODO docs
func (l FileLogger) Verbose(msg string) {
	l.Write(LogVerbose, msg)
}

// Info TODO docs
func (l FileLogger) Info(msg string) {
	l.Write(LogInfo, msg)
}

// Warning TODO docs
func (l FileLogger) Warning(msg string) {
	l.Write(LogWarning, msg)
}

// Error TODO docs
func (l FileLogger) Error(msg string) {
	l.Write(LogError, msg)
}

// Write TODO docs
func (l FileLogger) Write(lvl int, msg string) {
	if lvl < l.lvl {
		return
	}

	msg = fmt.Sprintf("[%s] %s\n", l.level(lvl), msg)

	if l.stdout {
		if color := l.color(lvl); color == "" {
			Print(msg)
		} else {
			PrintColor(color, msg)
		}
	}

	l.file.WriteString(msg)
}

func (l FileLogger) color(lvl int) string {
	switch lvl {
	case LogVerbose:
		return "lightblack+h"
	case LogInfo:
		return "cyan+h"
	case LogWarning:
		return "yellow+h"
	case LogError:
		return "red+h"
	}
	return ""
}

func (l FileLogger) level(lvl int) string {
	switch lvl {
	case LogVerbose:
		return "VERBOSE"
	case LogWarning:
		return "WARNING"
	case LogError:
		return "ERROR"
	case LogInfo:
		return "INFO"
	}
	return ""
}
