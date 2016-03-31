package fswatch

type Event struct {
	FilesModified []string
	FilesCreated  []string
	FilesDeleted  []string
}

func NewEvent(filesModified, filesCreated, filesDeleted []string) Event {
	return Event{
		FilesModified: filesModified,
		FilesCreated:  filesCreated,
		FilesDeleted:  filesDeleted,
	}
}
