package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

// Merge TODO docs
func Merge(ctx *cli.Context) error {
	yata.NewTaskManager().MergeFetchFiles()
	return nil
}
