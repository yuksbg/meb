package meb

import (
	"github.com/google/uuid"
	"sync"
)

type Event struct {
	Data interface{}
}

type Handler func(Event)

type EventBus struct {
	subscriptions map[string]map[string]Handler
	nextID        string
	mu            sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscriptions: make(map[string]map[string]Handler),
		nextID:        uuid.NewString(),
	}
}

func (eb *EventBus) Subscribe(eventType string, observer Handler) string {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if eb.subscriptions[eventType] == nil {
		eb.subscriptions[eventType] = make(map[string]Handler)
	}

	id := eb.nextID
	eb.nextID = uuid.NewString()

	eb.subscriptions[eventType][id] = observer
	return id
}

func (eb *EventBus) Unsubscribe(eventType string, id string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	delete(eb.subscriptions[eventType], id)
}

func (eb *EventBus) Publish(eventType string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	if observers, found := eb.subscriptions[eventType]; found {
		event := Event{Data: data}
		for _, observer := range observers {
			go observer(event)
		}
	}
}
