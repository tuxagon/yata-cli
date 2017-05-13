package cmd

import (
	"yata-cli/task"

	"github.com/urfave/cli"
)

// Prune removes any completed tasks
func Prune(ctx *cli.Context) error {
	m := task.NewFileManager()

	m.Prune()

	return nil
}
