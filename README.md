## fswatch

A command line tool and library package for watching file system changes.

### Installing the Command Line Tool and the Library

Install and build the binary.

```
go get github.com/wchan2/fswatch
```


### A Library Example

```go
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
```


## License

fswatch is licensed under the [MIT License](http://opensource.org/licenses/MIT).