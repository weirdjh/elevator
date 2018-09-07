package elevator

//https://github.com/kubernetes/apimachinery/blob/master/pkg/watch/watch.go

type EventType string

const (
	Added    EventType = "ADDED"
	Modified EventType = "MODIFIED"
	Deleted  EventType = "DELETED"
	Error    EventType = "ERROR"

	DefaultChanSize int32 = 100
)

type Event struct {
	ElevName string
	Type     EventType
	Msg      string
}

type EventWatch chan Event

func NewEventWatch() EventWatch {
	ch := make(chan Event)
	close(ch)
	return ch
}
