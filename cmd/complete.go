package cmd

import (
	"log"
	"strconv"

	"yata-cli/task"
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

	m := task.NewFileManager()
	task := m.GetTaskByID(uint32(id))
	if task != nil {
		task.Completed = true
		m.SaveTask(*task)
	} else {
		yata.PrintfColor("red+h", "Sorry, no task with an ID of %d was found", string(id))
	}

	return nil
}
