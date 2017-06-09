package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Fetch TODO docs
func Fetch(ctx *cli.Context) error {
	driveFetch := ctx.Bool("google-drive")

	if driveFetch {
		err := serverFetch(yata.GoogleDrive)
		if err != nil {
			return err
		}
	}

	return nil
}

func serverFetch(serverType int) error {
	manager := yata.NewServerManager(serverType)

	if err := manager.Fetch(); err != nil {
		yata.PrintlnColor("red+h", err.Error())
		return err
	}

	return nil
}
