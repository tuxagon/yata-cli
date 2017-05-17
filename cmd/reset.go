package cmd

import (
	"github.com/urfave/cli"
	//"github.com/tuxagon/yata-cli/task"
	"yata-cli/yata"
)

// Reset will erase any existing tasks and reset yata. By default,
// the old tasks will be backed up
func Reset(ctx *cli.Context) error {
	backup := !ctx.Bool("no-backup")
	resetID := !ctx.Bool("keep-id")

	manager := yata.NewTaskManager()

	if backup {
		manager.Backup()
	}
	return manager.Reset(resetID)
}
