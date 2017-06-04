package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

type pushArgs struct {
	googleDrive bool
}

func (a *pushArgs) Parse(ctx *cli.Context) {
	a.googleDrive = ctx.Bool("google-drive")
}

// Push TODO docs
func Push(ctx *cli.Context) error {
	args := &pushArgs{}
	args.Parse(ctx)

	if args.googleDrive {
		handleError(serverPush(yata.GoogleDrive))
	}

	return nil
}

func serverPush(serverType int) error {
	return yata.NewServerManager(serverType).Push()
}
