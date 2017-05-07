package yata

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

// FileManager handles all the task management through files
type FileManager struct {
	RootPath string
	FileName string
}

// NewFileManager creates a new file manager for yata tasks
func NewFileManager() *FileManager {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		panic(errors.New("Could not find 'HOME' environment variable"))
	}
	m := &FileManager{
		FileName: DefaultFilename,
		RootPath: path.Join(home, ".yata"),
	}
	m.Initialize()
	return m
}

// Initialize will create the .yata directory and all necessary files
// if it does not already exist. The default directory is '$HOME/.yata'
func (m *FileManager) Initialize() {
	_, err := os.Stat(m.RootPath)
	if err != nil {
		os.Mkdir(m.RootPath, 0777)
	}
	m.CreateDefaultTasksFile()
}

// CreateDefaultTasksFile will create the default task file in the root yata
// directory.
func (m *FileManager) CreateDefaultTasksFile() {
	m.CreateTasksFile("")
}

// CreateTasksFile will create a new file in the root for Yata tasks
// directory.
func (m *FileManager) CreateTasksFile(filename string) {
	fullPath := path.Join(m.RootPath, eitherString(filename, DefaultFilename))
	_, err := os.Stat(fullPath)
	if err != nil {
		ioutil.WriteFile(fullPath, []byte("[]"), 0777)
	}
}

// ReadFile reads the contents of the yata task file
func (m *FileManager) ReadFile() []byte {
	dat, err := ioutil.ReadFile(path.Join(m.RootPath, m.FileName))
	checkFatal(err)
	return dat
}

// GetAllOpenTasks gets any open task found in the yata file
func (m *FileManager) GetAllOpenTasks(s sort.Interface) (tasks []Task) {
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

// SaveNewTask will save the given task to the Yata file
func (m *FileManager) SaveNewTask(t Task) {
	tasks := m.GetAllOpenTasks(&ByPriority{})
	tasks = append(tasks, t)
	dat, err := json.Marshal(tasks)
	checkFatal(err)
	ioutil.WriteFile(path.Join(m.RootPath, m.FileName), dat, 0777)
}

// eitherString will return the first parameter if it is not nil; otherwise,
// it returns the second parameter
func eitherString(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	return s1
}
