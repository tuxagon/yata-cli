package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

type completeArgs struct {
	id int
}

func (a *completeArgs) Parse(ctx *cli.Context) {
	a.id = parseIDWithIndex(ctx, 0)
}

// Complete will mark the specified task as completed
func Complete(ctx *cli.Context) error {
	args := &completeArgs{}
	args.Parse(ctx)

	manager := yata.NewTaskManager()
	task, _ := manager.GetByID(uint32(args.id))
	if task != nil {
		task.Completed = true
		manager.Save(*task)
	} else {
		yata.PrintfColor("red+h", "Sorry, no task with an ID of %d was found", args.id)
	}

	return nil
}
