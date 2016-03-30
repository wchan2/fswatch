package fswatch

import (
	"io/ioutil"
	"log"
	"strings"
)

type FileLister interface {
	ListFiles() []string
}

type CompositeFileLister struct {
	filerListers []FileLister
}

func NewCompositeFileLister(filenames []string) FileLister {
	// TODO: need a smarter mechanism to know what is a subset of the others
	fileListers := make([]FileLister, len(filenames))
	for i, filename := range filenames {
		fileListers[i] = NewFileLister(filename)
	}

	return &CompositeFileLister{
		filerListers: fileListers,
	}
}

func (l *CompositeFileLister) ListFiles() (files []string) {
	for _, fileLister := range l.filerListers {
		files = append(files, fileLister.ListFiles()...)
	}
	return
}

func NewFileLister(fileExpression string) (fileLister FileLister) {
	if fileExpression == "**/*" {
		fileLister = NewRecursiveWildCardFileLister(".")
	} else if fileExpression == "*" {
		fileLister = NewCurrentDirFileLister()
	} else if strings.HasPrefix(fileExpression, "*") && strings.Count(fileExpression, "*") == 1 {
		fileLister = NewCurrentDirFileExtensionLister()
	} else if !strings.Contains(fileExpression, "*") {
		fileLister = NewNamedFileLister(fileExpression)
	}

	// TODO
	// other file listing cases
	// - **/*.mp4
	// - **/test.mp4
	// - test/test.mp4
	return
}

type RecursiveWildCardFileLister struct {
	directory string
}

func NewRecursiveWildCardFileLister(directory string) *RecursiveWildCardFileLister {
	return &RecursiveWildCardFileLister{directory: directory}
}

func (l *RecursiveWildCardFileLister) ListFiles() (returnFiles []string) {
	directories := []string{l.directory}
	for _, directory := range directories {
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			log.Printf("Unable to list files in %s", directory)
			continue
		}
		for _, file := range files {
			directories := []string{}
			if file.IsDir() {
				directories = append(directories, strings.Join([]string{directory, file.Name()}, "/"))
			} else {
				returnFiles = append(returnFiles, strings.Join([]string{directory, file.Name()}, "/"))
			}
		}
	}
	return
}

type CurrentDirFileLister struct{}

func NewCurrentDirFileLister() *CurrentDirFileLister {
	return &CurrentDirFileLister{}
}

func (l *CurrentDirFileLister) ListFiles() []string {
	// TODO: loop through only files in the current directory
	return []string{}
}

type CurrentDirFileExtensionLister struct{}

func NewCurrentDirFileExtensionLister() *CurrentDirFileExtensionLister {
	return &CurrentDirFileExtensionLister{}
}

func (l *CurrentDirFileExtensionLister) ListFiles() []string {
	// TODO: loop through the files that has a given suffix
	return []string{}
}

type NamedFileLister struct {
	filename string
}

func NewNamedFileLister(filename string) *NamedFileLister {
	return &NamedFileLister{filename: filename}
}

func (l *NamedFileLister) ListFiles() []string {
	// list exact matches for the file names (not nested, exact path)
	return []string{}
}
