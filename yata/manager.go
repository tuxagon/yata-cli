package yata

import "sort"

// TaskManager defines the behavior that should be found with any type of
// task management for yata
type TaskManager interface {
	Initialize()
	GetAllOpenTasks(s sort.Interface) []Task
	SaveNewTask(t Task)
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
