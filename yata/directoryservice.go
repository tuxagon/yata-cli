package yata

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const (
	configFilename    = ".yataconfig"
	defaultFilename   = "tasks"
	defaultPermission = 0777
	rootDirectory     = ".yata"
	idFilename        = ".yataid"
)
const defaultConfigContents = `Some config stuff needs to go here and a format`

var service *DirectoryService

// DirectoryService TODO docs
type DirectoryService struct {
	RootPath string
	Filename string
}

// NewDirectoryService TODO docs
func NewDirectoryService() *DirectoryService {
	if service == nil {
		home := getHomeDirectory()
		service = &DirectoryService{
			RootPath: filepath.Join(home, rootDirectory),
			Filename: defaultFilename,
		}
	}
	return service
}

// Initialize TODO docs
func (s DirectoryService) Initialize() {
	s.createRootPath()
	s.createTasksFile()
	s.createIDFile()
	s.createConfigFile()
}

// createRootPath TODO docs
func (s DirectoryService) createRootPath() {
	_, err := os.Stat(s.RootPath)
	if err != nil {
		os.Mkdir(s.RootPath, defaultPermission)
	}
}

// createTasksFile TODO docs
func (s DirectoryService) createTasksFile() {
	fullPath := s.getFullPath()
	if _, err := os.Stat(fullPath); err != nil {
		ioutil.WriteFile(fullPath, []byte("[]"), defaultPermission)
	}
}

// createIDFile TODO docs
func (s DirectoryService) createIDFile() {
	fullPath := s.getFullIDPath()
	if _, err := os.Stat(fullPath); err != nil {
		s.writeID(0)
	}
}

// createConfigFile TODO docs
func (s DirectoryService) createConfigFile() {
	fullPath := s.getFullConfigPath()
	if _, err := os.Stat(fullPath); err != nil {
		ioutil.WriteFile(fullPath, []byte(defaultConfigContents), defaultPermission)
	}
}

// getFullPath TODO docs
func (s DirectoryService) getFullPath() string {
	return path.Join(s.RootPath, s.Filename)
}

// getFullIDPath TODO docs
func (s DirectoryService) getFullIDPath() string {
	return path.Join(s.RootPath, idFilename)
}

// getFullConfigPath TODO docs
func (s DirectoryService) getFullConfigPath() string {
	return path.Join(s.RootPath, configFilename)
}

// writeID TODO docs
func (s DirectoryService) writeID(id uint32) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, id)
	ioutil.WriteFile(s.getFullIDPath(), bs, defaultPermission)
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
