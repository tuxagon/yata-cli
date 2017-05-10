package task

// Manager defines the behavior necessary to persist and manage
// tasks
type Manager interface {
	Initialize()
	GetAllOpenTasks() []Task
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
