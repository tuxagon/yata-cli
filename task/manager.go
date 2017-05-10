package task

// Manager defines the behavior necessary to persist and manage
// tasks
type Manager interface {
	Initialize()
	GetAllOpenTasks() []Task
	SaveNewTask(t Task)
}
