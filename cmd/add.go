package cmd

import (
	"yata-cli/task"

	//"github.com/tuxagon/yata-cli/task"
	"strings"

	"github.com/urfave/cli"
)

// Add creates a new task
func Add(ctx *cli.Context) error {
	m := task.NewFileManager()

	desc := ctx.String("description")
	if desc == "" {
		return cli.NewExitError("cannot add task without a description", 1)
	}

	tags := make([]string, 0)
	tagsFlag := ctx.String("tags")
	priority := ctx.Int("priority")

	if tagsFlag != "" {
		tags = strings.Split(tagsFlag, ",")
	}

	newTask := task.NewTask(desc, tags, priority)
	newTask.ExtractTags()

	m.SaveTask(*newTask)

	return nil
}
