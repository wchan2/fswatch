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
    "fmt"

    "github.com/wchan2/fswatch/fswatch"
)

func main() {
    fswatcher := fswatch.NewFileSystemWatcher(fswatch.DefaultHashFunc)
    fswatcher.Watch([]string{"text.txt"}).Subscribe(func(e fswatch.Event) {
        fmt.Println(e)
    })   
}
```


## License

fswatch is licensed under the [MIT License](http://opensource.org/licenses/MIT).