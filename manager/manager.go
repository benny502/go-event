package manager

import (
	"context"

	"github.com/benny502/go-event/listener"

	"github.com/benny502/go-event/event"

	"sync"
)

var once sync.Once

var instance *Manager

type Manager struct {
	lock      sync.Mutex
	container map[event.EventType][]listener.Listener
	eventChan chan event.Event
}

func (manager *Manager) Register(e event.EventType, l listener.Listener) {
	manager.lock.Lock()
	if _, ok := manager.container[e]; !ok {
		manager.container[e] = make([]listener.Listener, 0)
	}
	manager.container[e] = append(manager.container[e], l)
	manager.lock.Unlock()
}

func (manager *Manager) Send(e event.Event) {
	manager.eventChan <- e
}

func (manager *Manager) Start() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {

		for event := range manager.eventChan {
			select {
			case <-ctx.Done():
				return
			default:
				manager.lock.Lock()
				if value, ok := manager.container[event.Type]; ok {
					for _, listener := range value {
						go listener.Handler(event)
					}
				}
				manager.lock.Unlock()
			}
		}

	}()

	return cancel
}

func GetInstance() *Manager {

	once.Do(func() {
		instance = &Manager{
			lock:      sync.Mutex{},
			container: make(map[event.EventType][]listener.Listener),
			eventChan: make(chan event.Event, 10),
		}
	})
	return instance
}

func NewEvent(context context.Context, eventType event.EventType) event.Event {
	return event.Event{
		Type:    eventType,
		Value:   make(map[string]interface{}, 0),
		Context: context,
	}
}
