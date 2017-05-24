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

const (
	configFilename    = ".yataconfig.json"
	defaultFilename   = "tasks.json"
	defaultPermission = 0777
	rootDirectory     = ".yata"
	idFilename        = ".yataid"
)

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
func (s DirectoryService) Initialize() error {
	inits := []func() error{
		s.createRootPath,
		s.createTasksFile,
		s.createIDFile,
		s.createConfigFile,
	}
	for _, f := range inits {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

// GetCurrentID TODO docs
func (s DirectoryService) GetCurrentID() (id uint32, err error) {
	fullPath := s.GetFullIDPath()
	dat, err := ioutil.ReadFile(fullPath)
	return binary.BigEndian.Uint32(dat), err
}

// GetAndIncreaseID TODO docs
func (s DirectoryService) GetAndIncreaseID() (id uint32, err error) {
	currentID, err := s.GetCurrentID()
	if err != nil {
		return 0, err
	}

	newID := currentID + 1
	err = s.writeID(newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// GetConfig TODO docs
func (s DirectoryService) GetConfig() ([]byte, error) {
	fullPath := s.GetFullConfigPath()
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil, err
	}

	return ioutil.ReadFile(fullPath)
}

// WriteConfig TODO docs
func (s DirectoryService) WriteConfig(config []byte) error {
	return ioutil.WriteFile(s.GetFullConfigPath(), config, defaultPermission)
}

// Backup TODO docs
func (s DirectoryService) Backup() error {
	fullPath := s.GetFullPath()
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return err
	}

	dat, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return err
	}

	backupPath, err := s.GetBackupPath()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(backupPath, dat, defaultPermission)
}

// Reset TODO docs
func (s DirectoryService) Reset(resetID bool) error {
	fullPath := s.GetFullPath()
	if _, err := os.Stat(fullPath); err == nil {
		os.Remove(fullPath)
	}

	err := s.createTasksFile()
	if err != nil {
		return err
	}

	if resetID {
		return s.writeID(0)
	}
	return nil
}

// createRootPath TODO docs
func (s DirectoryService) createRootPath() error {
	_, err := os.Stat(s.RootPath)
	if err != nil {
		return os.Mkdir(s.RootPath, defaultPermission)
	}
	return nil
}

// createTasksFile TODO docs
func (s DirectoryService) createTasksFile() error {
	fullPath := s.GetFullPath()
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return ioutil.WriteFile(fullPath, []byte("[]"), defaultPermission)
	}
	return nil
}

// createIDFile TODO docs
func (s DirectoryService) createIDFile() error {
	fullPath := s.GetFullIDPath()
	if _, err := os.Stat(fullPath); err != nil {
		return s.writeID(0)
	}
	return nil
}

// createConfigFile TODO docs
func (s DirectoryService) createConfigFile() error {
	fullPath := s.GetFullConfigPath()
	if _, err := os.Stat(fullPath); err != nil {
		dat, _ := json.MarshalIndent(DefaultConfig(), "", "\t")
		return s.WriteConfig(dat)
	}
	return nil
}

// getFullPath TODO docs
func (s DirectoryService) GetFullPath() string {
	return filepath.Join(s.RootPath, s.Filename)
}

// getFullIDPath TODO docs
func (s DirectoryService) GetFullIDPath() string {
	return filepath.Join(s.RootPath, idFilename)
}

// getFullConfigPath TODO docs
func (s DirectoryService) GetFullConfigPath() string {
	return filepath.Join(s.RootPath, configFilename)
}

// getBackupPath TODO docs
func (s DirectoryService) GetBackupPath() (string, error) {
	files, err := ioutil.ReadDir(s.RootPath)
	if err != nil {
		return "", err
	}
	ext, n := ".bak", 0

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ext) {
			n++
		}
	}

	if n > 0 {
		ext = "." + string(n+int('0')) + ext
	}

	return s.GetFullPath() + ext, nil
}

// writeID TODO docs
func (s DirectoryService) writeID(id uint32) error {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, id)
	return ioutil.WriteFile(s.GetFullIDPath(), bs, defaultPermission)
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
