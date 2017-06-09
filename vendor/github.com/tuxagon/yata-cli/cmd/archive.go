package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Archive TODO docs
func Archive(ctx *cli.Context) error {
	archiver := yata.NewArchiver()
	err := archiver.Zip()
	if err != nil {
		yata.Println(err.Error())
		return err
	}
	return nil
}
