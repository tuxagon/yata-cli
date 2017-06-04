package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

// Prune removes any completed tasks
func Prune(ctx *cli.Context) error {
	handleError(yata.NewTaskManager().Prune())
	return nil
}
