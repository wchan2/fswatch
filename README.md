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
```

## License

fswatch is licensed under the [MIT License](http://opensource.org/licenses/MIT).