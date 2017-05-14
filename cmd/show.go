package cmd

import (
	"fmt"
	"log"
	"strconv"

	"yata-cli/task"
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// Show will display a task based on its ID
func Show(ctx *cli.Context) error {
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
		fmt.Println(task.String())
	} else {
		yata.PrintfColor("red+h", "Sorry, no task with an ID of %d was found", string(id))
	}

	return nil
}
