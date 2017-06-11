package yata

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Directory structure for Yata
//  .yata
//  |- md
//  |  |- id
//  |  |- fetch
//  |  |  |- id
//  |  |  |- tasks.json
//  |- .yataconfig
//  |- tasks.json
//

const (
	perm            = 0777
	rootDirName     = ".yata"
	metadataDirName = "md"
	fetchDirName    = "fetch"
	idFilename      = "id"
	configFilename  = ".yataconfig"
	tasksFilename   = "tasks.json"
	backupExt       = ".bak"
)

var dir *Directory

// Directory represents the Yata directory
type Directory struct {
	Root     string
	Filename string
}

// GetDirectory gets the Yata directory
func GetDirectory() *Directory {
	if dir == nil {
		home := getHomeDirectory()
		dir = &Directory{
			Root:     filepath.Join(home, rootDirName),
			Filename: tasksFilename,
		}
	}
	return dir
}

// Initialize will recreate any part of the Yata
// directory that is missing
func (d Directory) Initialize() {
	inits := []func() error{
		mkdir(d.Root),
		mkdir(d.MetadataDir()),
		mkdir(d.FetchDir()),
		d.createTasksFile,
		d.createIDFile,
		d.createConfigFile,
	}
	for _, f := range inits {
		HandleError(f(), true)
	}
}

// CurrentID returns the current ID in the Yata ID metadata file
func (d Directory) CurrentID() uint32 {
	dat, err := ioutil.ReadFile(d.IDPath())
	HandleError(err, true)
	return binary.BigEndian.Uint32(dat)
}

// IncrementID increments the current ID and returns the result
func (d Directory) IncrementID() uint32 {
	currentID := d.CurrentID()
	newID := currentID + 1
	d.WriteID(newID)
	return newID
}

// Config returns the entire config file as a byte slice
func (d Directory) Config() ([]byte, error) {
	path := d.ConfigPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	return ioutil.ReadFile(path)
}

// WriteConfig writes the entire config file at once
func (d Directory) WriteConfig(config []byte) error {
	return ioutil.WriteFile(d.ConfigPath(), config, perm)
}

// Backup creates a new backup of the current tasks file
func (d Directory) Backup() {
	path := d.TasksFilePath()
	_, err := os.Stat(path)
	HandleErrorWithFunc(err, true, func(err error) bool { return os.IsNotExist(err) })

	dat, err := ioutil.ReadFile(path)
	HandleError(err, true)

	backupPath := d.BackupPath()

	err = ioutil.WriteFile(backupPath, dat, perm)
	HandleError(err, true)
}

// Reset deletes current tasks and if resetID is true, will also
// set the ID back to 0
func (d Directory) Reset(resetID bool) {
	path := d.TasksFilePath()
	if _, err := os.Stat(path); err == nil {
		os.Remove(path)
	}

	d.createTasksFile()
	if resetID {
		d.WriteID(0)
	}
}

// EmptyFetchDir deletes any previously fetched files
func (d Directory) EmptyFetchDir() {
	os.RemoveAll(d.FetchDir())
	mkdir(d.FetchDir())()
}

// MetadataDir constructs the path to the metadata directory
func (d Directory) MetadataDir() string {
	return filepath.Join(d.Root, metadataDirName)
}

// FetchDir constructs the path to the fetch directory
func (d Directory) FetchDir() string {
	return filepath.Join(d.MetadataDir(), fetchDirName)
}

// TasksFilePath constructs the path to the tasks file
func (d Directory) TasksFilePath() string {
	return filepath.Join(d.Root, d.Filename)
}

// IDPath constructs the path to the ID file
func (d Directory) IDPath() string {
	return filepath.Join(d.MetadataDir(), idFilename)
}

// ConfigPath constructs the path to the config file
func (d Directory) ConfigPath() string {
	return filepath.Join(d.Root, configFilename)
}

// BackupPath constructs the path to the next valid backup path
func (d Directory) BackupPath() string {
	files, err := ioutil.ReadDir(d.Root)
	HandleError(err, true)
	ext, n := backupExt, 0

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ext) {
			n++
		}
	}

	if n > 0 {
		ext = "." + string(n+int('0')) + ext
	}

	return d.TasksFilePath() + ext
}

// WriteID saves the provided ID to the ID file
func (d Directory) WriteID(id uint32) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, id)
	err := ioutil.WriteFile(d.IDPath(), bs, perm)
	HandleError(err, true)
}

// createTasksFile creates the tasks file
func (d Directory) createTasksFile() error {
	path := d.TasksFilePath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ioutil.WriteFile(path, []byte("[]"), perm)
	}
	return nil
}

// createIDFile creates the ID file
func (d Directory) createIDFile() error {
	if _, err := os.Stat(d.IDPath()); err != nil {
		d.WriteID(0)
	}
	return nil
}

// createConfigFile creates the config file with defaults
func (d Directory) createConfigFile() error {
	if _, err := os.Stat(d.ConfigPath()); err != nil {
		dat, _ := json.MarshalIndent(DefaultConfig(), "", "\t")
		return d.WriteConfig(dat)
	}
	return nil
}

// mkdir creates a function that will create the specified directory
func mkdir(name string) func() error {
	return func() error {
		if _, err := os.Stat(name); err != nil {
			return os.Mkdir(name, perm)
		}
		return nil
	}
}

// getHomeDirectory will get the directory pointed to via the $HOME or %USERPROFILE%
// environment variable, depending on the OS
func getHomeDirectory() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	}
	home, ok := os.LookupEnv(env)
	if !ok {
		panic(fmt.Errorf("I could not find the '%s' environment variable", env))
	}
	return home
}
