package cmd

import (
	"sort"
	"strings"

	"yata-cli/yata"

	//"github.com/tuxagon/yata-cli/task"

	"github.com/urfave/cli"
)

// List returns the list of tasks/todos that have been recorded
func List(ctx *cli.Context) error {
	sort := ctx.String("sort")
	showAll := ctx.Bool("all")
	showTag := ctx.Bool("show-tags")
	searchTag := ctx.String("tag")
	searchDesc := ctx.String("description")
	format := ctx.String("format")

	manager := yata.NewTaskManager()
	tasks, err := manager.GetAll()
	if err != nil {
		return err
	}

	if showTag {
		return displayTags(tasks)
	}

	tasks = yata.FilterTasks(tasks, func(t yata.Task) bool {
		return (searchTag == "" || sliceContains(t.Tags, searchTag)) &&
			(searchDesc == "" || strings.Contains(t.Description, searchDesc)) &&
			(showAll || !t.Completed)
	})

	sortTasks(sort, &tasks)

	for _, v := range tasks {
		stringer := yata.NewTaskStringer(v, getStringerType(format))
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
	default:
		sort.Sort(yata.ByID(*tasks))
	}
}

func displayTags(tasks []yata.Task) error {
	tagCounts := make(map[string]int)
	for _, v := range tasks {
		for _, t := range v.Tags {
			_, ok := tagCounts[t]
			if !ok {
				tagCounts[t] = 1
			} else {
				tagCounts[t] = tagCounts[t] + 1
			}
		}
	}

	var tags []string
	maxLength := 0
	for k := range tagCounts {
		tags = append(tags, k)
		if len(k) > maxLength {
			maxLength = len(k)
		}
	}
	sort.Strings(tags)

	if len(tags) > 0 {

		for _, k := range tags {
			yata.Printf("%-*s\t%d\n", maxLength, k, tagCounts[k])
		}
	}
	return nil
}

func sliceContains(arr []string, term string) bool {
	for _, v := range arr {
		if v == term {
			return true
		}
	}

	return false
}

func getStringerType(format string) int8 {
	switch strings.ToLower(format) {
	case "json":
		return yata.JSON
	default:
		return yata.Simple
	}
}
