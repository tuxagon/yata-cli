package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

const (
	// Low represents a low priority
	Low = 1
	// Normal represents a normal, medium priority
	Normal = 2
	// High represents a high priority
	High = 3
)

const (
	// Simple represents a simple stringer
	Simple = iota
	// JSON represents a JSON stringer
	JSON
)

// Task represents a yata task
type Task struct {
	ID          uint32   `json:"id"`
	Description string   `json:"desc"`
	Completed   bool     `json:"done"`
	Priority    int      `json:"priority"`
	Tags        []string `json:"tags"`
	Timestamp   int64    `json:"timestamp"`
}

// SimpleStringer abstracts a simple task display
type SimpleStringer struct {
	task Task
}

// JSONStringer abstracts a task displaying with JSON
type JSONStringer struct {
	task Task
}

// Predicate is used to test some condition about a task
type Predicate func(Task) bool

// NewTask creates a new Yata task
func NewTask(description string, tags []string, priority int) *Task {
	return &Task{
		Description: description,
		Completed:   false,
		Priority:    priority,
		Tags:        tags,
		Timestamp:   time.Now().Unix(),
	}
}

// MarshalTask marshals the task into json bytes
func (t *Task) MarshalTask() []byte {
	json, err := json.Marshal(t)
	if err != nil {
		log.Fatal(errors.New("Unable to marshal task"))
	}
	return json
}

// ExtractTags will extract any tags from the description and
func (t *Task) ExtractTags() {
	re := regexp.MustCompile("#[A-z0-9_-]+")
	if tags := re.FindAllString(t.Description, -1); tags != nil {
		for i, v := range tags {
			tags[i] = v[1:]
		}
		t.Tags = append(t.Tags, tags...)
	}
}

// String returns a string representation of a Yata task
func (t *Task) String() string {
	dat, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return fmt.Sprintf("%+v", *t)
	}
	return string(dat)
}

// Filter returns a new list of tasks that satisfy the predicate
func Filter(tasks []Task, pred Predicate) []Task {
	filteredTasks := make([]Task, 0)
	for _, t := range tasks {
		if pred(t) {
			filteredTasks = append(filteredTasks, t)
		}
	}
	return filteredTasks
}

// NewTaskStringer returns the desired type that will stringify a
// task to display to the user
func NewTaskStringer(t Task, stringerType int8) fmt.Stringer {
	switch stringerType {
	case JSON:
		return JSONStringer{task: t}
	default:
		return SimpleStringer{task: t}
	}
}

func (s SimpleStringer) String() string {
	format := "{id} {description} {tags}"
	return replacePlaceholders(format, s.task)
}

func (s JSONStringer) String() string {
	dat, err := json.MarshalIndent(&s.task, "", "\t")
	if err != nil {
		return fmt.Sprintf("%+v", s.task)
	}
	return string(dat)
}

func replacePlaceholders(format string, t Task) string {
	if len(t.Tags) == 0 {
		format = strings.Replace(format, "{tags}", "", -1)
		format = strings.Trim(format, " ")
	}
	r := strings.NewReplacer(
		"{id}", fmt.Sprintf("%d", t.ID),
		"{description}", t.Description,
		"{priority}", fmt.Sprintf("%d", t.Priority),
		"{tags}", fmt.Sprintf("%+v", t.Tags),
		"{timestamp}", fmt.Sprintf("%d", t.Timestamp))
	return r.Replace(format)
}
