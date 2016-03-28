package main

import (
	"github.com/wchan2/fswatch"

	"log"
)

func main() {
	eventQ := make(chan fswatch.Event)
	fswatcher := fswatch.NewFileSystemWatcher(eventQ, []string{"text.txt"})
	go fswatcher.Start()
	defer fswatcher.Stop()

	for {
		select {
		case event := <-eventQ:
			log.Print(event)
		}
	}
}
