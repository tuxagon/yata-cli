package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

const (
	idPrompt = "Whoops, no task ID was specified. What is the ID of the task you want to complete?"
)

type cmdArgs interface {
	Parse(ctx *cli.Context)
}

type field string

func (f field) String() string { return string(f) }

func (f field) prompt() string {
	lf := strings.ToLower(f.String())
	yata.PrintfColor("yellow+h", "Missing %s: please provide a(n) %s\n%s: ", lf, lf, f.capitalize())
	return yata.Readln()
}

func (f field) promptInt() int {
	val := f.prompt()
	n, err := strconv.Atoi(val)
	if err != nil {
		showError(err.Error(), true)
	}
	return n
}

func (f field) capitalize() string {
	fs := f.String()
	if len(f) == 0 {
		return fs
	}
	return strings.ToUpper(string(fs[0])) + fs[1:]
}

func handleError(err error) {
	if err != nil {
		showError(err.Error(), true)
	}
}

func showError(msg string, exit bool) {
	yata.Println("red+h", msg)
	if exit {
		os.Exit(1)
	}
}

func taskStringer(format string) int8 {
	switch strings.ToLower(format) {
	case "json":
		return yata.JSON
	default:
		return yata.Simple
	}
}

func parseIDWithIndex(ctx *cli.Context, index int) (id int) {
	args := ctx.Args()

	id = ctx.Int("id")

	if id == 0 && len(args) == 0 {
		id = field("ID").promptInt()
	} else if id == 0 && len(args) > 0 {
		id, _ = strconv.Atoi(args[0])
	}

	if id == 0 {
		showError("No ID specified", true)
	}

	return
}
