package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Push TODO docs
func Push(ctx *cli.Context) error {
	manager := yata.NewServerManager(yata.GoogleDrive)

	if err := manager.Push(); err != nil {
		yata.PrintlnColor("red+h", err.Error())
		return err
	}

	return nil
}
