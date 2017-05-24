package main

import (
	"os"
	"sort"

	"yata-cli/cmd"
	"yata-cli/yata"

	//"github.com/tuxagon/yata-cli/cmd"

	"github.com/urfave/cli"
)

const (
	// Version specifies the current release
	Version = "1.0.0"

	descAdd      = "Create a new task"
	descArchive  = "Create an archive backup of the current tasks"
	descComplete = "Marks a task as done"
	descConfig   = "Manage configuration options"
	descDelete   = "Deletes a task"
	descList     = "Lists the tasks"
	descPrune    = "Removes all completed tasks"
	descPush     = "Uploads the tasks file"
	descReset    = "Erases all existing tasks and starts fresh"
	descShow     = "Displays a task based on its ID"
)

func main() {
	app := cli.NewApp()
	app.Name = "yata"
	app.Usage = "A command line task manager"
	app.Version = Version
	app.Before = func(ctx *cli.Context) error {
		dirService := yata.NewDirectoryService()
		if err := dirService.Initialize(); err != nil {
			yata.PrintlnColor("red+h", err.Error())
			os.Exit(1)
		}
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
		archive(),
		complete(),
		config(),
		delete(),
		list(),
		prune(),
		push(),
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
		Action:  cmd.Add,
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
	}
	sort.Sort(cli.FlagsByName(cmd.Flags))
	return cmd
}

func archive() cli.Command {
	cmd := cli.Command{
		Name:        "archive",
		Action:      cmd.Archive,
		Description: descArchive,
	}
	return cmd
}

func complete() cli.Command {
	cmd := cli.Command{
		Name:        "complete",
		Action:      cmd.Complete,
		Aliases:     []string{"finish"},
		Description: descComplete,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "specify the ID of the task to complete",
			},
		},
	}
	sort.Sort(cli.FlagsByName(cmd.Flags))
	return cmd
}

func config() cli.Command {
	return cli.Command{
		Name:        "config",
		Action:      cmd.Config,
		Description: descConfig,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "show-keys,k",
				Usage: "show the possible keys you can set/read",
			},
		},
	}
}

func delete() cli.Command {
	return cli.Command{
		Name:        "delete",
		Action:      cmd.Delete,
		Description: descDelete,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "specify the ID of the task to delete",
			},
		},
	}
}

func list() cli.Command {
	cmd := cli.Command{
		Name:   "list",
		Action: cmd.List,
		Usage:  descList,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "sort",
				Usage: "sort the results by the specified field",
			},
			cli.BoolFlag{
				Name:  "all,a",
				Usage: "display all tasks, including completed",
			},
		},
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

func push() cli.Command {
	cmd := cli.Command{
		Name:   "push",
		Usage:  descPush,
		Action: cmd.Push,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "google-drive,googledrive,drive,g",
				Usage: "uploads tasks file to Google Drive if you have an API key set in the config",
			},
		},
	}
	return cmd
}

func reset() cli.Command {
	cmd := cli.Command{
		Name:        "reset",
		Action:      cmd.Reset,
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
	}
	sort.Sort(cli.FlagsByName(cmd.Flags))
	return cmd
}

func show() cli.Command {
	cmd := cli.Command{
		Name:        "show",
		Action:      cmd.Show,
		Aliases:     []string{"get"},
		Description: descShow,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "specify the ID of the task to complete",
			},
		},
	}
	sort.Sort(cli.FlagsByName(cmd.Flags))
	return cmd
}
