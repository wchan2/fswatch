package main

import "github.com/wchan2/fswatch/fswatch"

func main() {
	fswatcher := fswatch.NewFileSystemWatcher(fswatch.DefaultHashFunc)
	fswatcher.Watch([]string{"text.txt"}).Subscribe(func(e fswatch.Event) {
		// do something
	})
}
