package yata

import (
	"fmt"
	"os"
	"path/filepath"
)

const logFile = ".yatalog"
const (
	LogVerbose = iota
	LogInfo
	LogWarning
	LogError
	LogNone
)

// LogLevel TODO docs
var LogLevel = LogVerbose

// Logger TODO docs
type Logger struct {
	stdout bool
	file   *os.File
	lvl    int
}

var logger *Logger

// GetLogger TODO docs
func GetLogger() *Logger {
	if logger != nil {
		return logger
	}

	filename := filepath.Join(NewDirectoryService().RootPath, logFile)

	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	logger = &Logger{
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
func (l Logger) Verbose(msg string) {
	l.Write(LogVerbose, msg)
}

// Info TODO docs
func (l Logger) Info(msg string) {
	l.Write(LogInfo, msg)
}

// Warning TODO docs
func (l Logger) Warning(msg string) {
	l.Write(LogWarning, msg)
}

// Error TODO docs
func (l Logger) Error(msg string) {
	l.Write(LogError, msg)
}

// Write TODO docs
func (l Logger) Write(lvl int, msg string) {
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

func (l Logger) color(lvl int) string {
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

func (l Logger) level(lvl int) string {
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
