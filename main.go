package main

import (
	"os"

	"github.com/tuxagon/yata-cli/cmd"
	"github.com/tuxagon/yata-cli/debug"
	"github.com/urfave/cli"
)

const (
	// Version is the current release of the command line app
	Version = "0.1.0"

	descList = "Lists the tasks/todos"
	descYata = "A command line task/todo manager"
)

func main() {
	app := cli.NewApp()
	app.Name = "yata"
	app.Usage = descYata
	app.Version = Version
	app.Before = func(ctx *cli.Context) error {
		debug.Verbose = ctx.GlobalBool("verbose")
		debug.Println("verbose logging enabled")

		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "turn on verbose logging",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "list",
			Usage: descList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "sort",
					Usage: "sort the results by the specified field",
				},
			},
			Action: cmd.List,
		},
	}
	app.Run(os.Args)
}
