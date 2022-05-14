package event

import "context"

type EventType int

type Event struct {
	Type    EventType
	Value   map[string]interface{}
	Context context.Context
}

func (event *Event) Set(key string, value interface{}) {
	event.Value[key] = value
}

func (event *Event) Get(key string) interface{} {
	value, ok := event.Value[key]
	if ok {
		return value
	}
	return nil
}
