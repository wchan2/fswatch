package fswatch

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type FileSystemWatcher struct {
	filenames []string
	eventQ    chan Event
	cache     *fileSystemCache
}

func NewFileSystemWatcher(filenames []string) *FileSystemWatcher {
	watcher := &FileSystemWatcher{
		filenames: filenames,
		eventQ:    make(chan Event),
		cache:     newFileSystemCache(),
	}
	watcher.cache.Set(watcher.fileHashes())
	return watcher
}

func (f *FileSystemWatcher) Start() {
	for {
		newFileHashes := f.fileHashes()
		if !f.cache.Equals(newFileHashes) {
			log.Printf("changed twice?", newFileHashes)
			filesChanged := f.cache.Diff(newFileHashes)
			events := map[string][]string{
				"changed": filesChanged,
				// "created": filesCreated,
			}
			f.cache.Set(newFileHashes)
			f.eventQ <- NewEvent(FileChange, events)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (f *FileSystemWatcher) Subscribe(callback func(Event)) {
	go func() {
		for event := range f.eventQ {
			callback(event)
		}
	}()
}

func (f *FileSystemWatcher) Stop() {
	close(f.eventQ)
}

func (f *FileSystemWatcher) fileHashes() map[string]string {
	fileHashes := map[string]string{}
	for _, filename := range f.filenames {
		fileContents, err := ioutil.ReadFile(filename)

		if err != nil {
			panic(err.Error())
		}
		fileHashes[filename] = f.hash(fileContents)
		if err != nil {
			panic(err.Error())
		}
	}
	return fileHashes
}

func (f *FileSystemWatcher) hash(contents []byte) string {
	return fmt.Sprintf("%x", md5.Sum(contents))
}
