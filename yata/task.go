package yata

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
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

// Manager handles all the task management
type Manager struct {
	RootPath string
	FileName string
}

// Task represents a yata task
type Task struct {
	Description string `json:"desc"`
	Completed   bool   `json:"done"`
	Priority    int    `json:"priority"`
	Project     string `json:"project"`
}

// ByPriority implements sort.Interface for []Task based on the Priority field
type ByPriority []Task

// ByDescription implements sort.Interface for []Task based on the Description field
type ByDescription []Task

func (t ByPriority) Len() int           { return len(t) }
func (t ByPriority) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByPriority) Less(i, j int) bool { return t[i].Priority < t[j].Priority }

func (t ByDescription) Len() int           { return len(t) }
func (t ByDescription) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByDescription) Less(i, j int) bool { return t[i].Description < t[j].Description }

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

// NewManager creates a new Yata manager
func NewManager() *Manager {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		panic(errors.New(ErrorPrefix + "Could not find 'HOME' environment variable"))
	}
	return &Manager{
		FileName: DefaultFilename,
		RootPath: path.Join(home, ".yata"),
	}
}

// InitializeDirectory will create the .yata directory and all necessary files
// if it does not already exist. The default directory is '$HOME/.yata'
func (m *Manager) InitializeDirectory() {
	_, err := os.Stat(m.RootPath)
	if err != nil {
		os.Mkdir(m.RootPath, 0777)
	}
	m.CreateDefaultTasksFile()
}

// CreateDefaultTasksFile will create the default task file in the root yata
// directory.
func (m *Manager) CreateDefaultTasksFile() {
	m.CreateTasksFile("")
}

// CreateTasksFile will create a new file in the root for Yata tasks
// directory.
func (m *Manager) CreateTasksFile(filename string) {
	fullPath := path.Join(m.RootPath, eitherString(filename, DefaultFilename))
	_, err := os.Stat(fullPath)
	if err != nil {
		ioutil.WriteFile(fullPath, []byte("[]"), 0777)
	}
}

// GetAllOpenTasks gets any open task found in the yata file
func (m *Manager) GetAllOpenTasks() (tasks []Task) {
	var rawTasks []Task
	dat := m.ReadFile()
	err := json.Unmarshal(dat, &rawTasks)
	checkFatal(err)
	for _, t := range rawTasks {
		if !t.Completed {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

// ReadFile reads the contents of the yata task file
func (m *Manager) ReadFile() []byte {
	dat, err := ioutil.ReadFile(path.Join(m.RootPath, m.FileName))
	checkFatal(err)
	return dat
}

// SaveNewTask will save the given task to the Yata file
func (m *Manager) SaveNewTask(t Task) {
	tasks := m.GetAllOpenTasks()
	tasks = append(tasks, t)
	dat, err := json.Marshal(tasks)
	checkFatal(err)
	ioutil.WriteFile(path.Join(m.RootPath, m.FileName), dat, 0777)
}

// eitherString will return the first parameter if it is not nil; otherwise, it returns the second parameter
func eitherString(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	return s1
}

// checkFatal displays a fatal log if
func checkFatal(err error) {
	if err != nil {
		log.Fatal(ErrorPrefix + err.Error())
	}
}
