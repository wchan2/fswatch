## fswatch [![Build Status](https://travis-ci.org/wchan2/fswatch.svg?branch=master)](https://travis-ci.org/wchan2/fswatch)

A library package for watching file system changes.

### Installation

```
go get github.com/wchan2/fswatch
```

### Example

```go
package main

import (
	"github.com/wchan2/fswatch"

	"log"
)

func main() {
	eventQ := make(chan fswatch.Event)
	fswatcher := fswatch.NewFileSystemWatcher(
		[]string{"**/*"},
		eventQ,
	)

	go fswatcher.Run()
	defer fswatcher.Stop()

	for {
		select {
		case event := <-eventQ:
			log.Print(event)
		}
	}
}
```

## TODO

A list of tasks to be completed

- [ ] Benchmark fswatch when there are a lot of files, large and small

## License

fswatch is licensed under the [MIT License](http://opensource.org/licenses/MIT).