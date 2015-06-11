package fswatch

type EventChannel <-chan Event
type Subscriber func(Event)

func (e EventChannel) Subscribe(callback Subscriber) {
	for event := range e {
		callback(event)
	}
}

type Event interface {
	Name() string
	Data() []byte
}

type changeEvent struct {
	name string
	data []byte
}

func NewEvent(name string, data []byte) Event {
	return changeEvent{
		name: name,
		data: data,
	}
}

func (c changeEvent) Name() string {
	return c.name
}

func (c changeEvent) Data() []byte {
	return c.data
}
