package cmd

import (
	"log"
	"strconv"

	"yata-cli/yata"

	"github.com/urfave/cli"
)

const (
	completeIDPrompt = "Whoops, no task ID was specified. What is the ID of the task you want to complete?"
)

// Complete will mark the specified task as completed
func Complete(ctx *cli.Context) error {
	args := ctx.Args()

	id := ctx.Int("id")

	if id == 0 && len(args) == 0 {
		yata.PrintfColor("yellow+h", "%s\nID: ", completeIDPrompt)
		id = yata.ReadInt()
	} else if id == 0 && len(args) > 0 {
		var err error
		id, err = strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err.Error())
		}
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
