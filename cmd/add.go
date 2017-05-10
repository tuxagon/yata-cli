package cmd

import (
	"yata-cli/task"

	//"github.com/tuxagon/yata-cli/task"
	"github.com/urfave/cli"
)

// Add creates a new task
func Add(ctx *cli.Context) error {
	m := task.NewFileManager()

	desc := ctx.String("description")
	proj := ctx.String("project")
	priority := ctx.Int("priority")

	newTask := task.NewTask(desc, proj, priority)

	m.SaveNewTask(*newTask)

	return nil
}
