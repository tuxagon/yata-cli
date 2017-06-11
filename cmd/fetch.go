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

// Fetch downloads any Yata files found on the specified server
func Fetch(ctx *cli.Context) error {
	args := &fetchArgs{}
	args.Parse(ctx)

	if args.googleDrive {
		doFetch(yata.GoogleDrive)
	}

	return nil
}

func doFetch(serverType int) {
	yata.NewSyncAPI(serverType).Fetch()
}
