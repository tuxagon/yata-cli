package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Delete Deletes a task
func Delete(ctx *cli.Context) error {
	id, err := promptForID(ctx)
	if err != nil {
		yata.PrintlnColor("red+h", err.Error())
	}

	manager := yata.NewTaskManager()

	return manager.DeleteByID(uint32(id))
}
