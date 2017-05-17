package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Prune removes any completed tasks
func Prune(ctx *cli.Context) error {
	manager := yata.NewTaskManager()
	return manager.Prune()
}
