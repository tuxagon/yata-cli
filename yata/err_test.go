package yata

import (
	"fmt"
	"testing"
)

type EmptyLogger struct {
	hit bool
}

func (l EmptyLogger) Verbose(msg string)        {}
func (l EmptyLogger) Info(msg string)           {}
func (l EmptyLogger) Warning(msg string)        {}
func (l *EmptyLogger) Error(msg string)         { l.hit = true }
func (l EmptyLogger) Write(lvl int, msg string) {}

func TestNilErrorDoesNotLog(t *testing.T) {
	l := mockLogger()

	HandleError(nil, false)

	if l.hit {
		t.Error("expected no logs for nil error")
	}
}

func TestNonNilErrorLogs(t *testing.T) {
	l := mockLogger()

	HandleError(fmt.Errorf(""), false)

	if !l.hit {
		t.Error("expected log for non-nil error")
	}
}

func TestNilErrorWithFuncDoesNotLog(t *testing.T) {
	l := mockLogger()

	HandleErrorWithFunc(nil, false, func(err error) bool { return false })

	if l.hit {
		t.Error("expected no logs for nil error with function")
	}
}

func TestNonNilErrorWithFuncDoesNotLog(t *testing.T) {
	l := mockLogger()

	HandleErrorWithFunc(fmt.Errorf(""), false, func(err error) bool { return false })

	if l.hit {
		t.Error("expected no logs for unsatisfied function")
	}
}

func TestNonNilErrorWithFuncLogs(t *testing.T) {
	l := mockLogger()

	HandleErrorWithFunc(fmt.Errorf(""), false, func(err error) bool { return true })

	if !l.hit {
		t.Error("expected log for non-nil error with function")
	}
}

func TestNilErrorDoesNotExit(t *testing.T) {
	mockLogger()
	HandleError(nil, true)
}

func mockLogger() *EmptyLogger {
	l := &EmptyLogger{}
	logger = l
	return l
}
