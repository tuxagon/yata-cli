package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
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

type YataConfig struct {
	BaseDir string
}

// YataManager handles all the task management
type YataManager struct {
	RootPath string
	FileName string
}

// YataTask represents a yata task
type YataTask struct {
	Description string `json:"desc"`
	Completed   bool   `json:"done"`
	Priority    int    `json:"priority"`
	Project     string `json:"project"`
}

// ByPriority implements sort.Interface for []YataTask based on the Priority field
type ByPriority []YataTask

func (t ByPriority) Len() int           { return len(t) }
func (t ByPriority) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByPriority) Less(i, j int) bool { return t[i].Priority < t[j].Priority }

// NewSimpleYataTask creates a new Yata task with defaults
func NewSimpleYataTask(description, project string) *YataTask {
	return &YataTask{
		Description: description,
		Completed:   false,
		Priority:    Med,
		Project:     EitherString(project, DefaultProject),
	}
}

// NewLowPriorityYataTask creates a new low priority Yata task with defaults
func NewLowPriorityYataTask(description, project string) *YataTask {
	return &YataTask{
		Description: description,
		Completed:   false,
		Priority:    Low,
		Project:     EitherString(project, DefaultProject),
	}
}

// NewHighPriorityYataTask creates a new high priority Yata task with defaults
func NewHighPriorityYataTask(description, project string) *YataTask {
	return &YataTask{
		Description: description,
		Completed:   false,
		Priority:    High,
		Project:     EitherString(project, DefaultProject),
	}
}

// MarshalTask marshals the task into json bytes
func (t YataTask) MarshalTask() []byte {
	json, err := json.Marshal(t)
	if err != nil {
		log.Fatal(errors.New("Unable to marshal task"))
	}
	return json
}

// String returns a string representation of a Yata task
func (t YataTask) String() string {
	return fmt.Sprintf("%s", t.Description)
}

// NewYataManager creates a new Yata manager
func NewYataManager() *YataManager {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		panic(errors.New(ErrorPrefix + "Could not find 'HOME' environment variable"))
	}
	return &YataManager{
		FileName: DefaultFilename,
		RootPath: path.Join(home, ".yata"),
	}
}

// InitializeDirectory will create the .yata directory if it does not already exist
func (m *YataManager) InitializeDirectory() {
	_, err := os.Stat(m.RootPath)
	if err != nil {
		os.Mkdir(m.RootPath, 0777)
	}
	m.CreateDefaultYataTasksFile()
}

// CreateDefaultYataTasksFile will create the default task file in the root
func (m *YataManager) CreateDefaultYataTasksFile() {
	m.CreateYataTasksFile("")
}

// CreateYataTasksFile will create a new file in the root for Yata tasks
func (m *YataManager) CreateYataTasksFile(filename string) {
	fullPath := path.Join(m.RootPath, EitherString(filename, DefaultFilename))
	_, err := os.Stat(fullPath)
	if err != nil {
		ioutil.WriteFile(fullPath, []byte("[]"), 0777)
	}
}

// GetAllOpenTasks gets any open task found in the yata file
func (m *YataManager) GetAllOpenTasks() (tasks []YataTask) {
	var rawTasks []YataTask
	dat := m.ReadFile()
	err := json.Unmarshal(dat, &rawTasks)
	CheckFatal(err)
	for _, t := range rawTasks {
		if !t.Completed {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

// ReadFile reads the contents of the yata task file
func (m *YataManager) ReadFile() []byte {
	dat, err := ioutil.ReadFile(path.Join(m.RootPath, m.FileName))
	CheckFatal(err)
	return dat
}

// SaveNewTask will save the given task to the Yata file
func (m *YataManager) SaveNewTask(t YataTask) {
	tasks := m.GetAllOpenTasks()
	tasks = append(tasks, t)
	dat, err := json.Marshal(tasks)
	CheckFatal(err)
	ioutil.WriteFile(path.Join(m.RootPath, m.FileName), dat, 0777)
}

// CheckFatal displays a fatal log if
func CheckFatal(err error) {
	if err != nil {
		log.Fatal(ErrorPrefix + err.Error())
	}
}

// EitherString will return the first parameter if it is not nil; otherwise, it returns the second parameter
func EitherString(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	return s1
}

// YataCmd declares functions that are implemented by each Yata command
type YataCmd interface {
	ParseOpts(args []string)
	Handle()
}

// ListCmd represents all the information regarding a list command
type ListCmd struct {
	Manager *YataManager
}

// ParseOpts TODO
func (cmd *ListCmd) ParseOpts(args []string) {

}

// Handle TODO
func (cmd ListCmd) Handle() {
	tasks := cmd.Manager.GetAllOpenTasks()
	sort.Sort(ByPriority(tasks))
	for _, t := range tasks {
		fmt.Println(t)
	}
}

func SeedFile(m *YataManager) {
	t1 := NewHighPriorityYataTask("Task 1", "")
	t2 := NewLowPriorityYataTask("Task 2", "")
	t3 := NewSimpleYataTask("Task 3", "")
	m.SaveNewTask(*t1)
	m.SaveNewTask(*t2)
	m.SaveNewTask(*t3)
}

func main() {
	ym := NewYataManager()
	ym.InitializeDirectory()
	//SeedFile(ym)

	if len(os.Args) <= 1 {
		// TODO Show usage
		os.Exit(1)
	}
	args := os.Args[1:]

	switch args[0] {
	case "add", "new":
		fmt.Println("New todo!")
	case "config":
		fmt.Println("Configuring")
	case "list", "ls":
		cmd := ListCmd{Manager: ym}
		cmd.Handle()
	default:
		fmt.Println("Usage")
	}

	fmt.Println(args)
}
