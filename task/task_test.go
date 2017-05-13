package task

import (
	"sort"
	"testing"
)

func TestFilterPredicateReturnsNoMatches(t *testing.T) {
	tasks := tasksOnlyId()
	filteredTasks := Filter(tasks, func(t Task) bool { return t.ID == 4 })

	if len(filteredTasks) > 0 {
		t.Errorf("expected [] got %+v", filteredTasks)
	}
}

func TestFilterPredicateReturnsOneMatch(t *testing.T) {
	tasks := tasksOnlyId()
	filteredTasks := Filter(tasks, func(t Task) bool { return t.ID == 1 })

	if len(filteredTasks) != 1 {
		t.Errorf("expected slice with 1 Task, got %+v", filteredTasks)
	}
}

func TestFilterPredicateReturnsMultipleMatches(t *testing.T) {
	tasks := tasksOnlyId()
	filteredTasks := Filter(tasks, func(t Task) bool { return t.ID > 1 })

	if len(filteredTasks) != 2 {
		t.Errorf("expected slice with 2 tasks, got %+v", filteredTasks)
	}
}

func TestNewTaskIsAlwaysMarkedIncomplete(t *testing.T) {
	task := NewTask("test", make([]string, 0), Normal)

	if task.Completed {
		t.Error("expected 'false' got 'true'")
	}
}

func TestMarshalTaskReturnsCorrectJson(t *testing.T) {
	task := NewTask("test", make([]string, 0), 2)
	json := string(task.MarshalTask())
	expected := "{\"id\":0,\"desc\":\"test\",\"done\":false,\"priority\":2,\"tags\":[]}"

	if json != expected {
		t.Errorf("expected %+v got %+v", expected, json)
	}
}

func TestExtractTagsWillExtractNoTagsFromDescriptionWithoutTags(t *testing.T) {
	task := NewTask("test", make([]string, 0), Normal)
	task.ExtractTags()

	if len(task.Tags) > 0 {
		t.Error("expected 0 got", len(task.Tags))
	}
}

func TestExtractTagsWillExtractAllTagsFromDescription(t *testing.T) {
	task := NewTask("help #frodo destroy #onering", make([]string, 0), Normal)
	task.ExtractTags()

	if len(task.Tags) != 2 {
		t.Error("expected 2 got", len(task.Tags))
		return
	}

	sort.Strings(task.Tags)

	if task.Tags[0] != "frodo" {
		t.Error("expected 'frodo' got", task.Tags[0])
	}

	if task.Tags[1] != "onering" {
		t.Error("expected 'onering' got", task.Tags[1])
	}
}

func TestExtractTagsWillPreserveExistingTags(t *testing.T) {
	task := NewTask("help #frodo", []string{"lotr"}, Normal)
	task.ExtractTags()

	if len(task.Tags) != 2 {
		t.Error("expected 2 got", len(task.Tags))
		return
	}

	sort.Strings(task.Tags)

	if task.Tags[1] != "lotr" {
		t.Error("expected 'lotr' got", task.Tags[1])
	}
}

func tasksOnlyId() []Task {
	return []Task{
		Task{ID: 1},
		Task{ID: 2},
		Task{ID: 3},
	}
}
