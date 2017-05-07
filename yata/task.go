package yata

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

const (
	// Low represents a low priority
	Low = 1 << iota
	// Med represents a medium priority
	Med
	// High represents a high priority
	High
)

const (
	// DefaultProject represents the name of the default project when on is not specified
	DefaultProject = "__none__"
	// DefaultFilename represents the name of the default Yata tasks file
	DefaultFilename = "main"
	// ErrorPrefix represents the string placed before an encountered error
	ErrorPrefix = "ERROR :: "
)

// Task represents a yata task
type Task struct {
	Description string `json:"desc"`
	Completed   bool   `json:"done"`
	Priority    int    `json:"priority"`
	Project     string `json:"project"`
}

// NewSimpleTask creates a new Yata task with defaults
func NewSimpleTask(description, project string) *Task {
	return NewTask(description, project, Med)
}

// NewLowPriorityTask creates a new low priority Yata task with defaults
func NewLowPriorityTask(description, project string) *Task {
	return NewTask(description, project, Low)
}

// NewHighPriorityTask creates a new high priority Yata task with defaults
func NewHighPriorityTask(description, project string) *Task {
	return NewTask(description, project, High)
}

// NewTask creates a new Yata task
func NewTask(description, project string, priority int) *Task {
	return &Task{
		Description: description,
		Completed:   false,
		Priority:    priority,
		Project:     eitherString(project, DefaultProject),
	}
}

// MarshalTask marshals the task into json bytes
func (t Task) MarshalTask() []byte {
	json, err := json.Marshal(t)
	if err != nil {
		log.Fatal(errors.New("Unable to marshal task"))
	}
	return json
}

// String returns a string representation of a Yata task
func (t Task) String() string {
	return fmt.Sprintf("%s", t.Description)
}

// checkFatal displays a fatal log if
func checkFatal(err error) {
	if err != nil {
		log.Fatal(ErrorPrefix + err.Error())
	}
}
