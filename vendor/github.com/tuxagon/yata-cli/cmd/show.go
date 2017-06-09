package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Show will display a task based on its ID
func Show(ctx *cli.Context) error {
	id, err := promptForID(ctx)
	if err != nil {
		yata.PrintlnColor("red+h", err.Error())
	}

	manager := yata.NewTaskManager()
	task, _ := manager.GetByID(uint32(id))

	if task != nil {
		stringer := yata.NewTaskStringer(*task, yata.Simple)
		yata.Println(stringer.String())
	} else {
		yata.PrintfColor("red+h", "Sorry, no task with an ID of %d was found\n", id)
	}

	return nil
}
