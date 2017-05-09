package task

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestDirectoryIsCreated(t *testing.T) {
	m := &FileManager{
		RootPath: ".yata",
		FileName: "task",
	}

	m.Initialize()
	if _, err := os.Stat(m.RootPath); err != nil {
		t.Error("expected directory '.yata' to exist got", err.Error())
	}

	os.RemoveAll(m.RootPath)
}

func TestTaskFileIsCreated(t *testing.T) {
	m := &FileManager{
		RootPath: ".yata",
		FileName: "task",
	}
	p := path.Join(m.RootPath, m.FileName)

	m.Initialize()
	if _, err := os.Stat(p); err != nil {
		t.Error("expected file '.yata/task' to exist got", err.Error())
	}

	os.RemoveAll(m.RootPath)
}

func TestTaskFileHasEmptyArray(t *testing.T) {
	m := &FileManager{
		RootPath: ".yata",
		FileName: "task",
	}
	p := path.Join(m.RootPath, m.FileName)

	m.Initialize()
	if dat, err := ioutil.ReadFile(p); err != nil {
		t.Error("expected file '.yata/task' to have contents '[]' got", err.Error())
	} else if string(dat) != "[]" {
		t.Error("expect file '.yata/task' to have contents '[]' got", dat)
	}

	os.RemoveAll(m.RootPath)
}
