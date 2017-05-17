package yata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	errMissingCollection = fmt.Errorf("missing collection")
)

// JSONDatabase TODO docs
type JSONDatabase struct {
	directory string
}

// NewJSONDatabase TODO docs
func NewJSONDatabase(dir string) Database {
	return &JSONDatabase{directory: dir}
}

// Read TODO docs
func (db JSONDatabase) Read(collection string, v interface{}) error {
	if collection == "" {
		return errMissingCollection
	}

	path := filepath.Join(db.directory, collection+".json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf(err.Error())
	}

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return json.Unmarshal(dat, &v)
}

func (db JSONDatabase) Write(collection string, v interface{}) error {
	if collection == "" {
		return errMissingCollection
	}

	path := filepath.Join(db.directory, collection+".json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf(err.Error())
	}

	dat, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, dat, 0777)
}
