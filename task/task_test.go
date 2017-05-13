package task_test

import (
	"testing"

	"yata-cli/task"
)

const (
	tasks = []task.Task{
		task.Task{ID: 1},
		task.Task{ID: 2},
		task.Task{ID: 3},
	}
)

func TestFilterResultsInEmptySlice(t *testing.T) {
	filteredTasks := tasks.Filter(func(t Task) { return t.ID == 4 })

	if len(filteredTasks) > 0 {
		t.Errorf("expected [] got %+v", filteredTasks)
	}
}

func TestFilterPredicateBringsSliceOfOne(t *testing.T) {
	filteredTasks := tasks.Filter(func(t Task) { return t.ID == 1 })

	if len(filteredTasks) != 1 {
		t.Errorf("expected slice with 1 Task, got %+v", filteredTasks)
	}
}

func TestFilterPredicateBringsBackMultiple(t *testing.T) {
	filteredTasks := tasks.Filter(func(t Task) { return t.ID > 1 })

	if len(filteredTasks) != 2 {
		t.Errorf("expected slice with 2 tasks, got %+v", filteredTasks)
	}
}
