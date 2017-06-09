package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Merge TODO docs
func Merge(ctx *cli.Context) error {
	manager := yata.NewTaskManager()
	err := manager.MergeFetchFiles()
	if err != nil {
		yata.PrintlnColor("red+h", err.Error())
	}
	return nil
}
