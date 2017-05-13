package cmd

import (
	"fmt"
	"strconv"
	"yata-cli/debug"
	"yata-cli/task"

	"github.com/urfave/cli"
)

// Complete will mark the specified task as completed
func Complete(ctx *cli.Context) error {
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
		task.Completed = true
		m.SaveTask(*task)
	} else {
		fmt.Printf("task %d not found\n", id)
	}

	return nil
}
