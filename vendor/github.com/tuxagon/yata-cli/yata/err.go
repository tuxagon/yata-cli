package yata

import (
	"os"
)

// ErrorCheckFunc represents a function that will be used in
// addition to the nil error check
//
// Example
//	func (err error) bool {
//		return err != io.EOF
//	}
type ErrorCheckFunc func(error) bool

// HandleError will log a non-nil error and if exit is true, then
// it will exit with an exit code of 1
func HandleError(err error, exit bool) {
	if err != nil {
		GetLogger().Error(err.Error())
		if exit {
			os.Exit(1)
		}
	}
}

// HandleErrorWithFunc will log a non-nil error if it also satisfies
// the provided function. If exit is true, it will exit with an exit
// code of 1
func HandleErrorWithFunc(err error, exit bool, fun ErrorCheckFunc) {
	if err != nil && fun(err) {
		GetLogger().Error(err.Error())
		if exit {
			os.Exit(1)
		}
	}
}
