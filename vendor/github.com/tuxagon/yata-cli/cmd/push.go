package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Push TODO docs
func Push(ctx *cli.Context) error {
	drivePush := ctx.Bool("google-drive")

	if drivePush {
		err := serverPush(yata.GoogleDrive)
		if err != nil {
			return err
		}
	}

	return nil
}

func serverPush(serverType int) error {
	manager := yata.NewServerManager(serverType)

	if err := manager.Push(); err != nil {
		yata.PrintlnColor("red+h", err.Error())
		return err
	}

	return nil
}
