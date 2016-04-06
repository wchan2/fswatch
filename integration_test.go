package fswatch_test

import (
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/wchan2/fswatch"
)

var _ = Describe(`FileSystemWatcher`, func() {
	Context(`when watching any file under any folder`, func() {
		var (
			fswatcher *fswatch.FileSystemWatcher
			eventQ    chan fswatch.Event
		)

		BeforeEach(func() {
			eventQ = make(chan fswatch.Event)
			fswatcher = fswatch.NewFileSystemWatcher(
				[]string{"**/*"},
				eventQ,
			)
		})

		JustBeforeEach(func() {
			go fswatcher.Run()

			if err := os.Mkdir("testing", os.ModePerm); err != nil {
				log.Fatal(err.Error())
			}
			if f, err := os.Create("testing/test_file.txt"); err != nil {
				log.Fatal(err.Error())
			} else {
				f.Chmod(os.ModePerm)
				f.Close()
			}
		})

		AfterEach(func() {
			fswatcher.Stop()

			if err := os.RemoveAll("testing"); err != nil {
				log.Fatal(err.Error())
			}
		})

		It(`receives an event for a created file`, func() {
			Expect(<-eventQ).To(Equal(fswatch.Event{
				FilesCreated: []string{"./testing/test_file.txt"},
			}))
		})
	})
})
