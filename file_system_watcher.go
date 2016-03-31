package fswatch

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type FileSystemWatcher struct {
	FileLister
	eventQ chan<- Event
	cache  StringCache
}

func NewFileSystemWatcher(filenames []string, eventQ chan<- Event) *FileSystemWatcher {
	watcher := &FileSystemWatcher{
		FileLister: NewCompositeFileLister(filenames),
		eventQ:     eventQ,
	}
	watcher.cache = watcher.fileHashes()
	return watcher
}

func (f *FileSystemWatcher) Start() {
	for {
		newHashes := f.fileHashes()
		if !f.cache.Equals(newHashes) {
			created, modified, deleted := f.cache.Diff(newHashes)
			f.cache = newHashes
			f.eventQ <- Event{
				FilesModified: modified,
				FilesCreated:  created,
				FilesDeleted:  deleted,
			}
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (f *FileSystemWatcher) Stop() {
	close(f.eventQ)
}

func (f *FileSystemWatcher) fileHashes() map[string]string {
	fileHashes := map[string]string{}
	for _, filename := range f.ListFiles() {
		fileContents, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Printf("Unable to open file %s because of %s", filename, err.Error())
			continue
		}
		fileHashes[filename] = f.hash(fileContents)
	}
	return fileHashes
}

func (f *FileSystemWatcher) hash(contents []byte) string {
	return fmt.Sprintf("%x", md5.Sum(contents))
}
