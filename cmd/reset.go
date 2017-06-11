package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

type resetArgs struct {
	noBackup bool
	keepID   bool
}

func (a *resetArgs) Parse(ctx *cli.Context) {
	a.noBackup = ctx.Bool("no-backup")
	a.keepID = ctx.Bool("keep-id")
}

// Reset will erase any existing tasks and reset yata. By default,
// the old tasks will be backed up
func Reset(ctx *cli.Context) error {
	args := &resetArgs{}
	args.Parse(ctx)

	manager := yata.NewTaskManager()

	if !args.noBackup {
		manager.Backup()
	}

	manager.Reset(!args.keepID)

	return nil
}
