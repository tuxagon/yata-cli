package main

import (
	"os"
	"time"
	"yata-cli/yata"

	"github.com/urfave/cli"
)

// func SeedFile(m *yata.Manager) {
// 	t1 := yata.NewHighPriorityTask("Task 1", "")
// 	t2 := yata.NewLowPriorityTask("Task 2", "")
// 	t3 := yata.NewSimpleTask("Task 3", "")
// 	m.SaveNewTask(*t1)
// 	m.SaveNewTask(*t2)
// 	m.SaveNewTask(*t3)
// }

func main() {
	app := cli.NewApp()
	app.Name = "yata"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Kenneth Bogner",
		},
	}
	app.Usage = "a command line task/todo manager"
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "list",
			ShortName:   "ls",
			Usage:       "[list] usage",
			Description: "[list] description",
		},
	}
	app.Run(os.Args)
	m := yata.NewFileManager()
	p := yata.NewCmdParser(m)
	if len(os.Args) > 1 {
		p.Parse(os.Args[1:])
	} else {
		p.Usage()
	}
}
