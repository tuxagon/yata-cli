package cmd

import (
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Complete will mark the specified task as completed
func Complete(ctx *cli.Context) error {
	id, err := promptForID(ctx)
	if err != nil {
		yata.PrintlnColor("red+h", err.Error())
	}

	manager := yata.NewTaskManager()
	task, _ := manager.GetByID(uint32(id))
	if task != nil {
		task.Completed = true
		manager.Save(*task)
	} else {
		yata.PrintfColor("red+h", "Sorry, no task with an ID of %d was found", string(id))
	}

	return nil
}
