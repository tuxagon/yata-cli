package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

type fetchArgs struct {
	googleDrive bool
}

func (a *fetchArgs) Parse(ctx *cli.Context) {
	a.googleDrive = ctx.Bool("google-drive")
}

// Fetch downloads any files applicable to yata from the specified server
func Fetch(ctx *cli.Context) error {
	args := &fetchArgs{}
	args.Parse(ctx)

	if args.googleDrive {
		handleError(serverFetch(yata.GoogleDrive))
	}

	return nil
}

func serverFetch(serverType int) error {
	return yata.NewServerManager(serverType).Fetch()
}
