package yata

import (
	"fmt"
	"time"
)

var (
	errTaskNotFound = fmt.Errorf("task not found")
)

// TaskManager TODO docs
type TaskManager struct {
	Database   Database
	Collection string
}

// NewTaskManager TODO docs
func NewTaskManager() *TaskManager {
	dirService := NewDirectoryService()
	return &TaskManager{
		Database:   NewDatabase(dirService.RootPath),
		Collection: "tasks",
	}
}

// GetAll TODO docs
func (m TaskManager) GetAll() (tasks []Task, err error) {
	tasks = make([]Task, 0)
	if err := m.Database.Read(m.Collection, &tasks); err != nil {
		return nil, err
	}
	return
}

// GetByID TODO docs
func (m TaskManager) GetByID(id uint32) (*Task, error) {
	tasks, err := m.GetAll()
	if err != nil {
		return nil, err
	}
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errTaskNotFound
}

// Save TODO docs
func (m TaskManager) Save(t Task) error {
	tasks, err := m.GetAll()
	if err != nil {
		return err
	}
	found := false
	t.Timestamp = time.Now().Unix()
	for i, v := range tasks {
		if v.ID == t.ID {
			changeIDIfZero(&t)
			tasks[i], found = t, true
			break
		}
	}

	if !found {
		changeIDIfZero(&t)
		tasks = append(tasks, t)
	}

	return m.Database.Write(m.Collection, tasks)
}

// DeleteByID TODO docs
func (m TaskManager) DeleteByID(id uint32) error {
	tasks, err := m.GetAll()
	if err != nil {
		return err
	}

	tasks = FilterTasks(tasks, func(t Task) bool {
		return t.ID != id
	})

	return m.Database.Write(m.Collection, tasks)
}

// Backup TODO docs
func (m TaskManager) Backup() error {
	dirService := NewDirectoryService()
	return dirService.Backup()
}

// Reset TODO docs
func (m TaskManager) Reset(resetID bool) error {
	dirService := NewDirectoryService()
	return dirService.Reset(resetID)
}

// Prune TODO docs
func (m TaskManager) Prune() error {
	tasks, err := m.GetAll()
	if err != nil {
		return err
	}

	tasks = FilterTasks(tasks, func(t Task) bool {
		return !t.Completed
	})

	return m.Database.Write(m.Collection, tasks)
}

// MergeFetchFiles TODO docs
func (m TaskManager) MergeFetchFiles() error {
	fetchDB := NewDatabase(NewDirectoryService().GetFetchPath())
	fetchTasks := make([]Task, 0)
	fetchDB.Read(m.Collection, &fetchTasks)
	tasks, err := m.GetAll()
	if err != nil {
		return err
	}
	maxID, err := NewDirectoryService().GetCurrentID()
	if err != nil {
		return err
	}

	hasUUID := func(tasks []Task, uuid string) bool {
		for _, t := range tasks {
			if t.UUID == uuid {
				return true
			}
		}
		return false
	}

	hasID := func(tasks []Task, id uint32) bool {
		for _, t := range tasks {
			if t.ID == id {
				return true
			}
		}
		return false
	}

	for _, ft := range fetchTasks {
		if hasID(tasks, ft.ID) {
			ft.ID, maxID = maxID+1, maxID+1
		}

		if !hasUUID(tasks, ft.UUID) {
			tasks = append(tasks, ft)
		}

		if maxID < ft.ID {
			maxID = ft.ID + 1
		}
	}

	NewDirectoryService().WriteID(maxID)

	return m.Database.Write(m.Collection, tasks)
}

// changeIDIfZero TODO docs
func changeIDIfZero(t *Task) error {
	dirService := NewDirectoryService()
	if t.ID == 0 {
		newID, err := dirService.GetAndIncreaseID()
		if err != nil {
			return err
		}
		t.ID = newID
	}
	return nil
}
