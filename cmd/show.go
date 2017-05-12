package cmd

import (
	"strconv"
	"yata-cli/debug"

	"yata-cli/task"

	"fmt"

	"github.com/urfave/cli"
)

// Show will display a task based on its ID
func Show(ctx *cli.Context) error {
	debug.Verbose = ctx.GlobalBool("verbose")
	debug.Printf("args :: %+v\n", ctx.Args())

	args := ctx.Args()

	if len(args) == 0 {
		return cli.NewExitError("a task id must be specified", 1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return cli.NewExitError("task id must be an integer", 1)
	}

	m := task.NewFileManager()
	task := m.GetTaskByID(uint32(id))
	if task != nil {
		fmt.Println(task.String())
	} else {
		fmt.Printf("task %d not found\n", id)
	}

	return nil
}
