package cmd

import (
	"sort"

	"yata-cli/task"
	"yata-cli/yata"

	//"github.com/tuxagon/yata-cli/task"

	"github.com/urfave/cli"
)

// List returns the list of tasks/todos that have been recorded
func List(ctx *cli.Context) error {
	var tasks []task.Task

	sort := ctx.String("sort")
	showAll := ctx.Bool("all")

	m := task.NewFileManager()
	if showAll {
		tasks = m.GetAllTasks()
	} else {
		tasks = m.GetAllOpenTasks()
	}

	tasks = sortTasks(sort, tasks)

	for _, v := range tasks {
		stringer := task.NewTaskStringer(v, task.Simple)
		switch v.Priority {
		case task.Low:
			yata.PrintlnColor("cyan+h", stringer.String())
		case task.High:
			yata.PrintlnColor("red+h", stringer.String())
		default:
			yata.Println(stringer.String())
		}
	}
	return nil
}

func sortTasks(sortField string, tasks []task.Task) []task.Task {
	switch {
	case sortField == "priority":
		sort.Sort(task.ByPriority(tasks))
	case sortField == "description":
		sort.Sort(task.ByDescription(tasks))
	case sortField == "timestamp":
		sort.Sort(task.ByTimestamp(tasks))
	case sortField != "":
		yata.PrintlnColor("yellow+h", "Sorry, but I can only sort using 'priority', 'description', or 'timestamp'. You should try one of those next time!")
	}
	return tasks
}
