package cmd

import (
	"fmt"

	"github.com/tuxagon/yata-cli/task"
	"github.com/urfave/cli"
)

// List returns the list of tasks/todos that have been recorded
func List(ctx *cli.Context) error {
	m := task.NewFileManager()
	tasks := m.GetAllOpenTasks()

	fmt.Println(tasks)
	for _, v := range tasks {
		fmt.Println(v)
	}
	return nil
}
