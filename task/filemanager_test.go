package task

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCreateRootPathWillCreateRootDirectory(t *testing.T) {
	m := newFileManager()

	m.createRootPath()

	if _, err := os.Stat(m.RootPath); err != nil {
		t.Error("expected directory '.yata' to exist got", err.Error())
	}

	cleanup(m)
}

func TestCreateTasksFileWillCreateTasksFile(t *testing.T) {
	m := newFileManager()
	m.createRootPath()
	p := path.Join(m.RootPath, m.FileName)

	m.createTasksFile()

	if _, err := os.Stat(p); err != nil {
		t.Error("expected file '.yata/task' to exist got", err.Error())
	}

	cleanup(m)
}

func TestCreateTasksFileWillWriteAnEmptySliceToFile(t *testing.T) {
	m := newFileManager()
	m.createRootPath()
	p := path.Join(m.RootPath, m.FileName)

	m.createTasksFile()

	if dat, _ := ioutil.ReadFile(p); string(dat) != "[]" {
		t.Error("expect file '.yata/task' to have contents '[]' got", string(dat))
	}

	cleanup(m)
}

func TestCreateTasksFileWillIgnoreTaskFileIfItExists(t *testing.T) {
	m := newFileManager()
	p := path.Join(m.RootPath, m.FileName)
	m.createRootPath()
	ioutil.WriteFile(p, []byte("test"), FilePermission)

	m.createTasksFile()

	if dat, _ := ioutil.ReadFile(p); string(dat) != "test" {
		t.Error("expect file '.yata/task' to have contents 'test' got", string(dat))
	}

	cleanup(m)
}

func TestCreateIDFileWillCreateIDFile(t *testing.T) {
	m := newFileManager()
	m.createRootPath()
	p := path.Join(m.RootPath, IDFilename)

	m.createIDFile()

	if _, err := ioutil.ReadFile(p); err != nil {
		t.Errorf("expected file '.yata/%s' to be read got %v", IDFilename, err.Error())
	}

	cleanup(m)
}

func TestCreateIDFileWillContainBigEndianZero(t *testing.T) {
	m := newFileManager()
	m.createRootPath()
	p := path.Join(m.RootPath, IDFilename)

	m.createIDFile()

	if dat, _ := ioutil.ReadFile(p); binary.BigEndian.Uint32(dat) != 0 {
		t.Error("expected file to contain 0 got", binary.BigEndian.Uint32(dat))
	}

	cleanup(m)
}

func TestCreateIDFileWillIgnoreTaskFileIfItExists(t *testing.T) {
	m := newFileManager()
	m.createRootPath()
	m.createIDFile()
	p := path.Join(m.RootPath, IDFilename)
	m.SetID(1)

	m.createIDFile()

	if dat, _ := ioutil.ReadFile(p); binary.BigEndian.Uint32(dat) != 1 {
		t.Error("expected file to contain 1 got", binary.BigEndian.Uint32(dat))
	}

	cleanup(m)
}

func TestGetFullPathShouldIncludeTheRootpathAndFilename(t *testing.T) {
	m := newFileManager()

	p := m.GetFullPath()

	if p != ".yata/task" {
		t.Error("expected '.yata/task' got", p)
	}
}

func TestSetIDWillSaveNumberToIDFile(t *testing.T) {
	m := newFileManager()
	m.Initialize()
	p := path.Join(m.RootPath, IDFilename)

	m.SetID(42)

	if dat, _ := ioutil.ReadFile(p); binary.BigEndian.Uint32(dat) != 42 {
		t.Error("expected file to contain 42 got", binary.BigEndian.Uint32(dat))
	}

	cleanup(m)
}

func TestReadFileWillReturnTheContentsOfTheFile(t *testing.T) {
	m := newFileManager()
	m.Initialize()

	dat := m.readFile()

	if string(dat) != "[]" {
		t.Error("expect file '.yata/task' to have contents '[]' got", string(dat))
	}

	cleanup(m)
}

func TestGetAllTasksWillReturnEveryTask(t *testing.T) {
	m := newFileManager()
	m.Initialize()
	tasks := mixedTasks()
	saveAllTasks(tasks, m)

	expectedTasks := m.GetAllTasks()

	if len(expectedTasks) != len(tasks) {
		t.Errorf("expected %d got %d", len(tasks), len(expectedTasks))
	}

	cleanup(m)
}

func TestGetAllOpenTasksWillReturnOnlyIncompleteTasks(t *testing.T) {
	m := newFileManager()
	m.Initialize()
	tasks := mixedTasks()
	saveAllTasks(tasks, m)

	expectedTasks := m.GetAllOpenTasks()

	if len(expectedTasks) != 2 {
		t.Errorf("expected 2 got %d", len(expectedTasks))
	}

	for _, v := range expectedTasks {
		if v.Completed {
			t.Error("expected only incomplete tasks, got", v)
		}
	}

	cleanup(m)
}

func TestGetTaskByIDWillReturnOnlyOneTaskWithThatID(t *testing.T) {
	m := newFileManager()
	m.Initialize()
	tasks := mixedTasks()
	saveAllTasks(tasks, m)

	task := m.GetTaskByID(2)

	if task.ID != 2 {
		t.Errorf("expected task with ID of 2 got %v", task)
	}

	cleanup(m)
}

func TestSaveTaskWillSaveANewTask(t *testing.T) {
	m := newFileManager()
	m.Initialize()
	newTask := NewTask("test", make([]string, 0), Normal)

	m.SaveTask(*newTask)

	if task := m.GetTaskByID(1); task == nil {
		t.Errorf("expected task to be saved got no new task")
	}

	cleanup(m)
}

func TestSaveTaskWillUpdateExisingTask(t *testing.T) {
	m := newFileManager()
	m.Initialize()
	tasks := mixedTasks()
	saveAllTasks(tasks, m)
	oldTask := m.GetTaskByID(1)
	oldTask.Description = "mordor"

	m.SaveTask(*oldTask)

	if allTasks := m.GetAllTasks(); len(tasks) != len(allTasks) {
		t.Errorf("expected %d got %d", len(tasks), len(allTasks))
	}

	if task := m.GetTaskByID(1); task == nil || task.Description != "mordor" {
		t.Error("expected 'mordor' got", task.Description)
	}

	cleanup(m)
}

func newFileManager() *FileManager {
	return &FileManager{
		RootPath: ".yata",
		FileName: "task",
	}
}

func mixedTasks() []Task {
	return []Task{
		Task{ID: 1, Completed: true},
		Task{ID: 2, Completed: false},
		Task{ID: 3, Completed: false},
	}
}

func saveAllTasks(tasks []Task, m *FileManager) {
	for _, t := range tasks {
		m.SaveTask(t)
	}
}

func cleanup(m *FileManager) {
	os.RemoveAll(m.RootPath)
}
