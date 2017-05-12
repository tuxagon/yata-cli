package task

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"

	"strings"

	"encoding/binary"
)

const (
	// DefaultProject represents the name of the default project when on is not specified
	DefaultProject = "__none__"
	// DefaultFilename represents the name of the default Yata tasks file
	DefaultFilename = "tasks"
	// IDFilename represents the name of the file containing the next task id to use
	IDFilename = ".yataid"
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
	m.CreateTasksFile()
	m.CreateIDFile()
}

// CreateTasksFile will create a new tasks file in the root for Yata tasks
// directory with the specified name
func (m *FileManager) CreateTasksFile() {
	fullPath := m.GetFullPath()
	if _, err := os.Stat(fullPath); err != nil {
		ioutil.WriteFile(fullPath, []byte("[]"), 0777)
	}
}

// ReadFile reads the contents of a yata task file
func (m *FileManager) ReadFile() []byte {
	dat, err := ioutil.ReadFile(m.GetFullPath())
	checkFatal(err)
	return dat
}

// GetAllOpenTasks gets any open task found in the yata file
func (m *FileManager) GetAllOpenTasks() (tasks []Task) {
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
	t.ID = m.GetAndIncreaseID()
	tasks := m.GetAllOpenTasks()
	tasks = append(tasks, t)
	dat, err := json.Marshal(tasks)
	checkFatal(err)
	ioutil.WriteFile(m.GetFullPath(), dat, 0777)
}

// BackUp will copy the tasks file into a separate file
func (m *FileManager) BackUp() {
	fullPath := m.GetFullPath()
	if _, err := os.Stat(fullPath); err != nil {
		return
	}
	dat, err := ioutil.ReadFile(fullPath)
	checkFatal(err)
	ioutil.WriteFile(m.GetBackUpFile(), dat, 0777)
}

// Reset will erase the existing tasks file
func (m *FileManager) Reset() {
	fullPath := m.GetFullPath()
	if _, err := os.Stat(fullPath); err == nil {
		os.Remove(fullPath)
	}
	m.CreateTasksFile()
	m.SetID(0)
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
	checkFatal(err)
	return binary.BigEndian.Uint32(dat)
}

// GetAndIncreaseID will return the ID from the ID file and increment the ID in the file
func (m *FileManager) GetAndIncreaseID() uint32 {
	newID := m.GetCurrentID() + 1
	m.SetID(newID)
	return newID
}

// CreateIDFile will create the yata ID file if it does not exist
func (m *FileManager) CreateIDFile() {
	fullPath := m.GetFullIDPath()
	if _, err := os.Stat(fullPath); err != nil {
		m.SetID(0)
	}
}

// SetID writes the id specified into the ID file
func (m *FileManager) SetID(id uint32) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, id)
	ioutil.WriteFile(m.GetFullIDPath(), bs, 0777)
}

// eitherString will return the first parameter if it is not nil; otherwise,
// it returns the second parameter
func eitherString(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	return s1
}

// checkFatal displays a fatal log if
func checkFatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
