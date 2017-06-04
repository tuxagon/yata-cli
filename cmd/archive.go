package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

// Archive creates an archive file of the current tasks
func Archive(ctx *cli.Context) error {
	archiver := yata.NewArchiver()
	err := archiver.Zip()
	if err != nil {
		showError(err.Error(), true)
	}
	return nil
}
