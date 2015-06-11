package fswatch

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func DefaultHashFunc(contents []byte) string {
	return fmt.Sprintf("%x", md5.Sum(contents))
}

type Watcher interface {
	Watch(filename []string) EventChannel
}

type watcher struct {
	output chan Event
	cache  FSCache
	hashFn func([]byte) string
}

func NewFileSystemWatcher(hashFn func([]byte) string) Watcher {
	return &watcher{
		output: make(chan Event),
		hashFn: hashFn,
		cache:  NewFileSystemCache(),
	}
}

func (w *watcher) Watch(filenames []string) EventChannel {
	w.cache.Update(w.calculateFileHashes(filenames))

	go func() {
		for {
			fileHashes := w.calculateFileHashes(filenames)
			if !w.cache.Equals(fileHashes) {
				filesChanged := w.cache.Diff(fileHashes)
				eventData, err := json.Marshal(map[string][]string{"changed": filesChanged})
				if err != nil {
					panic(err.Error())
				}

				w.cache.Update(fileHashes)
				w.publish(NewEvent("fileChange", eventData))
			}
		}
	}()

	return w.output
}

func (w *watcher) calculateFileHashes(filenames []string) map[string]string {
	fileHashes := map[string]string{}
	for _, filename := range filenames {
		fileContents, err := ioutil.ReadFile(filename)

		if err != nil {
			panic(err.Error())
		}
		fileHashes[filename] = w.hashFn(fileContents)
		if err != nil {
			panic(err.Error())
		}
	}
	return fileHashes
}

func (w *watcher) publish(event Event) {
	w.output <- event
}
