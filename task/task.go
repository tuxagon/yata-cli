package task

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

// Task represents a yata task
type Task struct {
	ID          string `json:"id"`
	Description string `json:"desc"`
	Completed   bool   `json:"done"`
	Priority    int    `json:"priority"`
	Project     string `json:"project"`
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
func (t *Task) MarshalTask() []byte {
	json, err := json.Marshal(t)
	if err != nil {
		log.Fatal(errors.New("Unable to marshal task"))
	}
	return json
}

// String returns a string representation of a Yata task
func (t *Task) String() string {
	return fmt.Sprintf("%s", t.Description)
}
