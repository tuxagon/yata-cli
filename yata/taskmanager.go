package yata

const (
	// RootDirectory TODO docs
	RootDirectory = ".yata"
)

type TaskManager struct {
	Dir DirectoryService
	Db  Database
}
