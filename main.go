package main

import (
	"os"
	"sort"
	"yata-cli/cmd"

	//"github.com/tuxagon/yata-cli/cmd"
	"github.com/tuxagon/yata-cli/debug"
	"github.com/urfave/cli"
)

const (
	// Version is the current release of the command line app
	Version = "0.1.0"

	descAdd  = "Adds a new task"
	descList = "Lists the tasks"
	descYata = "A command line task manager"
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
		cli.Command{
			Name:  "add",
			Usage: descAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "description,desc,d",
					Usage: "specify the task description",
				},
				cli.StringFlag{
					Name:  "project,proj",
					Usage: "specify the name of a project for a task",
				},
				cli.IntFlag{
					Name:  "priority,p",
					Usage: "specify a priority for the task (1: Low, 2: Normal, 3: High)",
				},
			},
			Action: cmd.Add,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
