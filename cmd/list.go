package cmd

import (
	"sort"

	"yata-cli/yata"

	//"github.com/tuxagon/yata-cli/task"

	"github.com/urfave/cli"
)

// List returns the list of tasks/todos that have been recorded
func List(ctx *cli.Context) error {
	sort := ctx.String("sort")
	showAll := ctx.Bool("all")

	manager := yata.NewTaskManager()
	tasks, err := manager.GetAll()
	if err != nil {
		return err
	}

	if !showAll {
		tasks = yata.FilterTasks(tasks, func(t yata.Task) bool {
			return !t.Completed
		})
	}

	sortTasks(sort, &tasks)

	for _, v := range tasks {
		stringer := yata.NewTaskStringer(v, yata.Simple)
		switch v.Priority {
		case yata.LowPriority:
			yata.PrintlnColor("cyan+h", stringer.String())
		case yata.HighPriority:
			yata.PrintlnColor("red+h", stringer.String())
		default:
			yata.Println(stringer.String())
		}
	}
	return nil
}

func sortTasks(sortField string, tasks *[]yata.Task) {
	switch {
	case sortField == "priority":
		sort.Sort(yata.ByPriority(*tasks))
	case sortField == "description":
		sort.Sort(yata.ByDescription(*tasks))
	case sortField == "timestamp":
		sort.Sort(yata.ByTimestamp(*tasks))
	case sortField != "":
		yata.PrintlnColor("yellow+h", "Sorry, but I can only sort using 'priority', 'description', or 'timestamp'. You should try one of those next time!")
	}
}
