package cmd

import (
	"fmt"
	"strconv"

	"yata-cli/yata"

	"github.com/urfave/cli"
)

const (
	idPrompt = "Whoops, no task ID was specified. What is the ID of the task you want to complete?"
)

func promptForID(ctx *cli.Context) (int, error) {
	return promptForIDWithIndex(ctx, 0)
}

func promptForIDWithIndex(ctx *cli.Context, idx int) (id int, err error) {
	args := ctx.Args()

	id = ctx.Int("id")

	if id == 0 && len(args) == 0 {
		yata.PrintlnColor("yellow+h", idPrompt)
		yata.Print("ID: ")
		return yata.ReadInt(), nil
	}

	if id == 0 && len(args) > idx {
		return strconv.Atoi(args[idx])
	}

	if id == 0 {
		return id, fmt.Errorf("Whoops, no task ID was specified")
	}

	return id, nil
}
