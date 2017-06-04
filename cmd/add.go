package cmd

import (
	"strings"

	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

type addArgs struct {
	tags        []string
	priority    int
	description string
}

func (a *addArgs) Parse(ctx *cli.Context) {
	args := ctx.Args()

	tags := ctx.String("tags")
	a.priority = ctx.Int("priority")
	a.description = ctx.String("description")

	if a.description == "" && len(args) == 0 {
		a.description = field("description").prompt()
	} else if a.description == "" && len(args) > 0 {
		a.description = strings.Join(args, " ")
	}

	if a.priority <= 0 || a.priority > 3 {
		a.priority = yata.NormalPriority
	}

	if tags != "" {
		a.tags = strings.Split(tags, ",")
	}
}

// Add creates a new task
func Add(ctx *cli.Context) error {
	args := &addArgs{}
	args.Parse(ctx)

	manager := yata.NewTaskManager()
	newTask := yata.NewTask(args.description, args.tags, args.priority)
	newTask.ExtractTags()

	return manager.Save(*newTask)
}
