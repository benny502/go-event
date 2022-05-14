package listener

import "github.com/benny502/go-event/event"

type Listener interface {
	Handler(event.Event)
}
