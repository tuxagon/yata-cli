package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
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
	ID          uint32   `json:"id"`
	Description string   `json:"desc"`
	Completed   bool     `json:"done"`
	Priority    int      `json:"priority"`
	Tags        []string `json:"tags"`
}

// NewTask creates a new Yata task
func NewTask(description string, tags []string, priority int) *Task {
	return &Task{
		Description: description,
		Completed:   false,
		Priority:    priority,
		Tags:        tags,
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

// ExtractTagsFromDescription will extract any tags from the description and
func (t *Task) ExtractTagsFromDescription() {
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
	return fmt.Sprintf("%s", t.Description)
}
