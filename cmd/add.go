package cmd

import (

	//"github.com/tuxagon/yata-cli/task"
	"strings"

	"yata-cli/yata"

	"github.com/urfave/cli"
)

const (
	addDescriptionPrompt = "Whoops, that task is missing a description. What description should this task have?"
)

// Add creates a new task
func Add(ctx *cli.Context) error {
	args := ctx.Args()

	tags := make([]string, 0)
	flagTags := ctx.String("tags")
	priority := ctx.Int("priority")
	description := ctx.String("description")

	if description == "" && len(args) == 0 {
		yata.PrintlnColor("yellow+h", addDescriptionPrompt)
		yata.Printf("Description: ")
		description = yata.Readln()
	} else if description == "" && len(args) > 0 {
		description = strings.Join(args, " ")
	}

	if priority == 0 || priority > 3 {
		priority = yata.NormalPriority
	}

	if flagTags != "" {
		tags = strings.Split(flagTags, ",")
	}

	manager := yata.NewTaskManager()
	newTask := yata.NewTask(description, tags, priority)
	newTask.ExtractTags()

	return manager.Save(*newTask)
}
