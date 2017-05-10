package cmd

import (
	"fmt"
	"sort"
	"yata-cli/task"

	//"github.com/tuxagon/yata-cli/task"
	"github.com/urfave/cli"
)

// List returns the list of tasks/todos that have been recorded
func List(ctx *cli.Context) error {
	m := task.NewFileManager()
	tasks := m.GetAllOpenTasks()

	tasks = sortTasks(ctx.String("sort"), tasks)

	for _, v := range tasks {
		fmt.Println(v.String())
	}
	return nil
}

func sortTasks(sortField string, tasks []task.Task) []task.Task {
	switch sortField {
	case "priority":
		sort.Sort(task.ByPriority(tasks))
	case "description":
		sort.Sort(task.ByDescription(tasks))
	}

	return tasks
}
