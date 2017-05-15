package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"strings"

	"encoding/binary"
	"runtime"
)

const (
	// DefaultFilename represents the name of the default Yata tasks file
	DefaultFilename = "tasks"
	// IDFilename represents the name of the file containing the next task id to use
	IDFilename = ".yataid"
	// FilePermission represents the file permissions each file should get
	FilePermission = 0777
)

// FileManager handles all the task management through files
type FileManager struct {
	RootPath string
	FileName string
}

// NewFileManager creates a new file manager for yata tasks
func NewFileManager() *FileManager {
	home := getHomeDirectory()
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
	m.createRootPath()
	m.createTasksFile()
	m.createIDFile()
}

// GetAllTasks get any task found in the yata file
func (m *FileManager) GetAllTasks() (tasks []Task) {
	dat := m.readFile()
	err := json.Unmarshal(dat, &tasks)
	exitIfErr(err)
	return
}

// GetAllOpenTasks gets any open task found in the yata file
func (m *FileManager) GetAllOpenTasks() (tasks []Task) {
	rawTasks := m.GetAllTasks()
	tasks = make([]Task, 0)
	for _, t := range rawTasks {
		if !t.Completed {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

// GetTaskByID returns a task by the specified ID
func (m *FileManager) GetTaskByID(id uint32) *Task {
	tasks := m.GetAllTasks()
	for _, t := range tasks {
		if t.ID == id {
			return &t
		}
	}
	return nil
}

// SaveTask will save the given task to the Yata file
func (m *FileManager) SaveTask(t Task) {
	tasks := m.GetAllTasks()
	found := false
	t.Timestamp = time.Now().Unix()
	for i, v := range tasks {
		if v.ID == t.ID {
			m.changeIDIfZero(&t)
			tasks[i] = t
			found = true
			break
		}
	}

	if !found {
		m.changeIDIfZero(&t)
		tasks = append(tasks, t)
	}

	dat, err := json.Marshal(tasks)
	exitIfErr(err)
	ioutil.WriteFile(m.GetFullPath(), dat, FilePermission)
}

// BackUp will copy the tasks file into a separate file
func (m *FileManager) BackUp() {
	fullPath := m.GetFullPath()
	if _, err := os.Stat(fullPath); err != nil {
		return
	}
	dat, err := ioutil.ReadFile(fullPath)
	exitIfErr(err)
	ioutil.WriteFile(m.GetBackUpFile(), dat, FilePermission)
}

// Reset will erase the existing tasks file
func (m *FileManager) Reset() {
	fullPath := m.GetFullPath()
	if _, err := os.Stat(fullPath); err == nil {
		os.Remove(fullPath)
	}
	m.createTasksFile()
	m.SetID(0)
}

// Prune removes any completed tasks from the tasks file
func (m *FileManager) Prune() {
	tasks := m.GetAllTasks()
	fullPath := m.GetFullPath()
	ioutil.WriteFile(fullPath, []byte("[]"), FilePermission)
	for _, t := range tasks {
		if !t.Completed {
			m.SaveTask(t)
		}
	}
}

// GetFilename gets the name of the tasks file from config
func (m *FileManager) GetFilename() string {
	return m.FileName
}

// GetFullPath returns the full path to the tasks file
func (m *FileManager) GetFullPath() string {
	return path.Join(m.RootPath, m.GetFilename())
}

// GetFullIDPath returns the full path to the yata ID file
func (m *FileManager) GetFullIDPath() string {
	return path.Join(m.RootPath, IDFilename)
}

// GetBackUpFile returns the name of the back up file
func (m *FileManager) GetBackUpFile() string {
	files, _ := ioutil.ReadDir(m.RootPath)
	ext := ".bak"
	n := 0
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".bak") {
			n++
		}
	}
	if n > 0 {
		ext = "." + string(n+int('0')) + ext
	}
	return m.GetFullPath() + ext
}

// GetCurrentID gets the current ID from the ID file
func (m *FileManager) GetCurrentID() uint32 {
	fullPath := m.GetFullIDPath()
	dat, err := ioutil.ReadFile(fullPath)
	exitIfErr(err)
	return binary.BigEndian.Uint32(dat)
}

// GetAndIncreaseID will return the ID from the ID file and increment the ID in the file
func (m *FileManager) GetAndIncreaseID() uint32 {
	newID := m.GetCurrentID() + 1
	m.SetID(newID)
	return newID
}

// SetID writes the id specified into the ID file
func (m *FileManager) SetID(id uint32) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, id)
	ioutil.WriteFile(m.GetFullIDPath(), bs, FilePermission)
}

// createRootPath will create the root directory for tasks if it does
// not already exist
func (m FileManager) createRootPath() {
	_, err := os.Stat(m.RootPath)
	if err != nil {
		os.Mkdir(m.RootPath, FilePermission)
	}
}

// createTasksFile will create a new tasks file in the root of the
// tasks directory
func (m FileManager) createTasksFile() {
	fullPath := m.GetFullPath()
	if _, err := os.Stat(fullPath); err != nil {
		ioutil.WriteFile(fullPath, []byte("[]"), FilePermission)
	}
}

// createIDFile will create the .yataid file if it does not exist
func (m FileManager) createIDFile() {
	fullPath := m.GetFullIDPath()
	if _, err := os.Stat(fullPath); err != nil {
		m.SetID(0)
	}
}

// readFile reads the contents of a yata task file
func (m FileManager) readFile() []byte {
	dat, err := ioutil.ReadFile(m.GetFullPath())
	exitIfErr(err)
	return dat
}

func (m FileManager) changeIDIfZero(t *Task) {
	if t.ID == 0 {
		t.ID = m.GetAndIncreaseID()
	}
}

// eitherString will return the first parameter if it is not nil; otherwise,
// it returns the second parameter
func eitherString(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	return s1
}

// exitIfErr displays a fatal log if
func exitIfErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// getHomeDirectory will get the directory pointed to via the $HOME or %USERPROFILE%
// environment variable, depending on the OS
func getHomeDirectory() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	}
	home, ok := os.LookupEnv(env)
	if !ok {
		panic(fmt.Errorf("Could not find '%s' environment variable", env))
	}
	return home
}
