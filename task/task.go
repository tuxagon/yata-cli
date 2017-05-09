package task

import (
	"encoding/json"
	"errors"
	"log"
)

// Task represents a yata task
type Task struct {
	Description string `json:"desc"`
	Completed   bool   `json:"done"`
	Priority    int    `json:"priority"`
	Project     string `json:"project"`
}


// MarshalTask marshals the task into json bytes
func (t *Task) MarshalTask() []byte {
	json, err := json.MarshalTask(t)
	if err != nil {
		log.Fatal(errors.New("unable to marshal task"))
	}
	return json
}