package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

type showArgs struct {
	id     int
	format string
}

func (a *showArgs) Parse(ctx *cli.Context) {
	a.id = parseIDWithIndex(ctx, 0)
	a.format = ctx.String("format")
}

// Show will display a task based on its ID
func Show(ctx *cli.Context) error {
	args := &showArgs{}
	args.Parse(ctx)

	manager := yata.NewTaskManager()
	task, _ := manager.GetByID(uint32(args.id))

	if task != nil {
		stringer := yata.NewTaskStringer(*task, taskStringer(args.format))
		yata.Println(stringer.String())
	} else {
		yata.PrintfColor("red+h", "Sorry, no task with an ID of %d was found\n", args.id)
	}

	return nil
}
