package yata

import (
	"io"
	"os"
	"path/filepath"
)

const logFile = ".yatalog"

// Logger TODO docs
type Logger struct {
	outs []io.Writer
}

var logger *Logger

// NewLogger TODO docs
func NewLogger() *Logger {
	if logger != nil {
		return logger
	}

	fp, err := os.Open(filepath.Join(NewDirectoryService().RootPath, logFile))
	logger = &Logger{
		outs: []io.Writer{
			os.Stdout,
			fp,
		},
	}

	if err != nil {
		logger.Warning(err.Error())
	}

	return logger
}

// Warning TODO docs
func (l Logger) Warning(msg string) {
	msg = "[WARNING] " + msg
	PrintlnColor("yellow+h", msg)
}
