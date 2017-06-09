package cmd

import (
	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

type deleteArgs struct {
	id int
}

func (a *deleteArgs) Parse(ctx *cli.Context) {
	a.id = parseIDWithIndex(ctx, 0)
}

// Delete deletes a task
func Delete(ctx *cli.Context) error {
	args := &deleteArgs{}
	args.Parse(ctx)

	manager := yata.NewTaskManager()
	return manager.DeleteByID(uint32(args.id))
}
