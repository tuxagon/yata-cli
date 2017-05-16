package yata

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

// JSONDatabase TODO docs
type JSONDatabase struct {
	Database *scribble.Driver
}

// NewJSONDatabase TODO docs
func NewJSONDatabase(dir string) Database {
	db, err := scribble.New(dir, nil)
	if err != nil {
		PrintlnColor("red+h", err.Error())
	}
	return &JSONDatabase{Database: db}
}

// Read TODO docs
func (db *JSONDatabase) Read(v interface{}) error {
	return db.Database.Read("tasks", "tasks", &v)
}

func (db *JSONDatabase) Placeholder() {

}
