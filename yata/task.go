package yata

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

const (
	LowPriority    = 1
	NormalPriority = 2
	HighPriority   = 3
)

const (
	Simple = iota
	JSON
)

// Task TODO docs
type Task struct {
	ID          uint32   `json:"id"`
	Description string   `json:"desc"`
	Completed   bool     `json:"done"`
	Priority    int      `json:"priority"`
	Tags        []string `json:"tags"`
	Timestamp   int64    `json:"timestamp"`
	UUID        string   `json:"uuid"`
}

// SimpleStringer TODO docs
type SimpleStringer struct {
	task Task
}

// JSONStringer TODO docs
type JSONStringer struct {
	task Task
}

// Predicate TODO docs
type Predicate func(Task) bool

// ByPriority TODO docs
type ByPriority []Task

// ByDescription TODO docs
type ByDescription []Task

// ByTimestamp TODO docs
type ByTimestamp []Task

func (t ByPriority) Len() int           { return len(t) }
func (t ByPriority) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByPriority) Less(i, j int) bool { return t[i].Priority < t[j].Priority }

func (t ByDescription) Len() int           { return len(t) }
func (t ByDescription) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByDescription) Less(i, j int) bool { return t[i].Description < t[j].Description }

func (t ByTimestamp) Len() int           { return len(t) }
func (t ByTimestamp) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByTimestamp) Less(i, j int) bool { return t[i].Timestamp < t[j].Timestamp }

// NewTask TODO docs
func NewTask(description string, tags []string, priority int) *Task {
	uuid, _ := uuid.NewV4()
	return &Task{
		Description: description,
		Completed:   false,
		Priority:    priority,
		Tags:        tags,
		Timestamp:   time.Now().Unix(),
		UUID:        uuid.String(),
	}
}

// ExtractTags TODO docs
func (t *Task) ExtractTags() {
	re := regexp.MustCompile("#[A-z0-9_-]+")
	if tags := re.FindAllString(t.Description, -1); tags != nil {
		for i, v := range tags {
			tags[i] = v[1:]
		}
		t.Tags = append(t.Tags, tags...)
	}
}

// MarshalTask TODO docs
func (t *Task) MarshalTask() (dat []byte, err error) {
	return json.Marshal(t)
}

// String TODO docs
func (t *Task) String() string {
	dat, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return fmt.Sprintf("%+v", *t)
	}
	return string(dat)
}

// FilterTasks TODO docs
func FilterTasks(tasks []Task, pred Predicate) []Task {
	filteredTasks := make([]Task, 0)
	for _, t := range tasks {
		if pred(t) {
			filteredTasks = append(filteredTasks, t)
		}
	}
	return filteredTasks
}

// NewTaskStringer TODO docs
func NewTaskStringer(t Task, stringerType int8) fmt.Stringer {
	switch stringerType {
	case JSON:
		return JSONStringer{task: t}
	default:
		return SimpleStringer{task: t}
	}
}

// String TODO docs
func (s SimpleStringer) String() string {
	format := "{id} {description} {tags}"
	return replacePlaceholders(format, s.task)
}

// String TODO docs
func (s JSONStringer) String() string {
	dat, err := json.MarshalIndent(&s.task, "", "\t")
	if err != nil {
		return fmt.Sprintf("%+v", s.task)
	}
	return string(dat)
}

// replacePlaceholders TODO docs
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
		"{timestamp}", fmt.Sprintf("%d", t.Timestamp),
		"{uuid}", t.UUID)
	return r.Replace(format)
}
