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

		// TODO: iterate over current hashes and figure out if a file is changed, created, deleted to propagate the correct event
		if !f.cache.Equals(newHashes) {
			filesChanged := f.cache.Diff(newHashes)
			f.cache = newHashes
			f.eventQ <- NewEvent(FileChanged, map[string][]string{
				"changed": filesChanged,
			})
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
