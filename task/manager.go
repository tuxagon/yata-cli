package task

// TaskManager defines the behavior necessary to persist and manage
// tasks
type TaskManager interface {
	Initialize()
	GetAllOpenTasks() []Task
	SaveNewTask(t Task)
}