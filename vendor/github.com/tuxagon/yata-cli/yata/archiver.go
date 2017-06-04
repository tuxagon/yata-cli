package yata

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const archiveFilename = "yata.tasks"

// Archiver TODO docs
type Archiver struct{}

// archivableFile TODO docs
type archivableFile struct {
	Filename string
	Contents []byte
}

// NewArchiver TODO docs
func NewArchiver() *Archiver {
	return &Archiver{}
}

// Zip TODO docs
func (a Archiver) Zip() error {
	dirService := NewDirectoryService()
	fw, err := os.Create(filepath.Join(dirService.RootPath, archiveFilename+".zip"))
	w := zip.NewWriter(fw)
	defer w.Close()

	files, err := a.getArchivableFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		f, err := w.Create(file.Filename)
		if err != nil {
			return err
		}
		_, err = f.Write(file.Contents)
		if err != nil {
			return err
		}
	}

	return nil
}

// getArchivableFiles TODO docs
func (a Archiver) getArchivableFiles() ([]archivableFile, error) {
	dirService := NewDirectoryService()
	files := make([]archivableFile, 0)
	filenames := []string{
		dirService.GetFullPath(),
		dirService.GetFullIDPath(),
	}

	for _, f := range filenames {
		dat, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}
		files = append(files, archivableFile{
			Filename: stripPath(dirService.RootPath+"/", f),
			Contents: dat,
		})
	}

	return files, nil
}

func stripPath(pathToStrip, fullPath string) string {
	return strings.Replace(fullPath, pathToStrip, "", -1)
}
