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

// Push uploads the Yata files needed to synchronize tasks across
// many machines
func Push(ctx *cli.Context) error {
	args := &pushArgs{}
	args.Parse(ctx)

	if args.googleDrive {
		doPush(yata.GoogleDrive)
	}

	return nil
}

func doPush(serverType int) {
	yata.NewSyncAPI(serverType).Push()
}
