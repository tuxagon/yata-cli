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
	// Version specifies the current release
	Version = "1.0.0"

	descAdd      = "Create a new task"
	descComplete = "Marks a task as done"
	descConfig   = "Manage configuration options"
	descList     = "Lists the tasks"
	descPrune    = "Removes all completed tasks"
	descReset    = "Erases all existing tasks and starts fresh"
	descShow     = "Displays a task based on its ID"
)

func main() {
	app := cli.NewApp()
	app.Name = "yata"
	app.Usage = "A command line task manager"
	app.Version = Version
	app.Before = func(ctx *cli.Context) error {
		debug.Verbose = ctx.GlobalBool("verbose")
		debug.Println("info :: verbose logging enabled")
		return nil
	}
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Kenneth Bogner",
		},
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "turn on verbose logging",
		},
	}
	app.Commands = []cli.Command{
		add(),
		complete(),
		list(),
		prune(),
		reset(),
		show(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}

func add() cli.Command {
	cmd := cli.Command{
		Name:    "add",
		Aliases: []string{"new", "create"},
		Usage:   descAdd,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "description,desc,d",
				Usage: "specify the task description; tags can be included in description using the '#' prefix in the description text",
			},
			cli.StringFlag{
				Name:  "tags,t",
				Usage: "specify tags outside of the description; list is comma-delimited",
			},
			cli.IntFlag{
				Name:  "priority,p",
				Usage: "specify a priority for the task (1: Low, 2: Normal, 3: High)",
			},
		},
		Action: cmd.Add,
	}
	sort.Sort(cli.FlagsByName(cmd.Flags))
	return cmd
}

func complete() cli.Command {
	cmd := cli.Command{
		Name:        "complete",
		Aliases:     []string{"finish"},
		Description: descComplete,
		Action:      cmd.Complete,
	}
	sort.Sort(cli.FlagsByName(cmd.Flags))
	return cmd
}

func config() cli.Command {
	return cli.Command{
		Name:        "config",
		Description: descConfig,
		Action:      cmd.Config,
	}
}

func list() cli.Command {
	cmd := cli.Command{
		Name:  "list",
		Usage: descList,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "sort",
				Usage: "sort the results by the specified field",
			},
		},
		Action: cmd.List,
	}
	sort.Sort(cli.FlagsByName(cmd.Flags))
	return cmd
}

func prune() cli.Command {
	cmd := cli.Command{
		Name:   "prune",
		Usage:  descPrune,
		Action: cmd.Prune,
	}
	return cmd
}

func reset() cli.Command {
	cmd := cli.Command{
		Name:        "reset",
		Description: descReset,
		Aliases:     []string{"nuke"},
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "no-backup",
				Usage: "prevent yata from making a backup before resetting",
			},
			cli.BoolFlag{
				Name:  "keep-id",
				Usage: "keep the current incrementing ID rather than starting from 1 again",
			},
		},
		Action: cmd.Reset,
	}
	sort.Sort(cli.FlagsByName(cmd.Flags))
	return cmd
}

func show() cli.Command {
	cmd := cli.Command{
		Name:        "show",
		Aliases:     []string{"get"},
		Description: descShow,
		Action:      cmd.Show,
	}
	return cmd
}
