package fswatch

const FileChange = "FileChange"

type Event struct {
	Type string
	Data map[string][]string
}

func NewEvent(eventType string, data map[string][]string) Event {
	return Event{
		Type: eventType,
		Data: data,
	}
}
