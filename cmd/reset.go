package cmd

import (
	"yata-cli/task"

	"github.com/urfave/cli"
	//"github.com/tuxagon/yata-cli/task"
)

// Reset will erase any existing tasks and reset yata. By default,
// the old tasks will be backed up
func Reset(ctx *cli.Context) error {
	m := task.NewFileManager()

	if !ctx.Bool("no-backup") {
		m.BackUp()
	}
	if !ctx.Bool("keep-id") {
		m.SetID(0)
	}
	m.Reset()

	return nil
}
