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
	fw, err := os.Create(filepath.Join(GetDirectory().Root, archiveFilename+".zip"))
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
	files := make([]archivableFile, 0)
	filenames := []string{
		GetDirectory().TasksFilePath(),
		GetDirectory().IDPath(),
	}

	for _, f := range filenames {
		dat, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}
		files = append(files, archivableFile{
			Filename: stripPath(GetDirectory().Root+"/", f),
			Contents: dat,
		})
	}

	return files, nil
}

func stripPath(pathToStrip, fullPath string) string {
	return strings.Replace(fullPath, pathToStrip, "", -1)
}
