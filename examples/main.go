package main

import (
	"log"

	"github.com/wchan2/fswatch"
)

func main() {
	fswatcher := fswatch.NewFileSystemWatcher([]string{"text.txt"})
	fswatcher.Subscribe(func(event fswatch.Event) {
		log.Print(event)
	})
	fswatcher.Start()
}
