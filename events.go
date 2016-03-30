package fswatch

const (
	FileCreated FileEvent = "CREATED"
	FileChanged FileEvent = "CHANGED"
	FileDeleted FileEvent = "DELETED"
)

type FileEvent string

type Event struct {
	Type FileEvent
	Data map[string][]string
}

func NewEvent(eventType FileEvent, data map[string][]string) Event {
	return Event{
		Type: eventType,
		Data: data,
	}
}
