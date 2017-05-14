package cmd

import (
	"yata-cli/task"

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
		yata.PrintfColor("yellow+h", "%s\nDescription: ", addDescriptionPrompt)
		description = yata.Readln()
	} else if description == "" && len(args) > 0 {
		description = strings.Join(args, " ")
	}

	m := task.NewFileManager()

	if flagTags != "" {
		tags = strings.Split(flagTags, ",")
	}

	newTask := task.NewTask(description, tags, priority)
	newTask.ExtractTags()

	m.SaveTask(*newTask)

	return nil
}
