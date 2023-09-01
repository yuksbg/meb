package meb

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
	var receivedEvents []Event
	eb := NewEventBus()
	eb.Subscribe("testEvent", func(e Event) {
		receivedEvents = append(receivedEvents, e)
	})
	eb.Publish("testEvent", "testData")

	time.Sleep(time.Millisecond * 10)

	assert.Equal(t, 1, len(receivedEvents), "Events number != 1")
	assert.Equal(t, "testData", receivedEvents[0].Data, "Event data is not correct")
}

func TestUnsubscribe(t *testing.T) {
	var receivedEvents []Event
	eb := NewEventBus()
	id := eb.Subscribe("testEvent", func(e Event) {
		receivedEvents = append(receivedEvents, e)
	})

	eb.Publish("testEvent", "testData1")
	eb.Unsubscribe("testEvent", id)
	eb.Publish("testEvent", "testData2")
	time.Sleep(time.Millisecond * 10)

	assert.Equal(t, 1, len(receivedEvents), "Events number != 1")
	assert.Equal(t, "testData1", receivedEvents[0].Data, "Event data is not correct")
}

func TestMultipleSubscribe(t *testing.T) {
	eb := NewEventBus()

	var wg sync.WaitGroup
	receivedEvens1 := make(chan Event, 1)
	receivedEvens2 := make(chan Event, 1)

	wg.Add(2)

	id1 := eb.Subscribe("testEvent", func(e Event) { receivedEvens1 <- e; wg.Done() })
	id2 := eb.Subscribe("testEvent", func(e Event) { receivedEvens2 <- e; wg.Done() })

	time.Sleep(10 * time.Millisecond) // just in case?!

	eb.Publish("testEvent", "testData")
	wg.Wait()

	eb.Unsubscribe("testEvent", id1)
	eb.Unsubscribe("testEvent", id2)

	e1 := <-receivedEvens1
	e2 := <-receivedEvens2

	assert.Equal(t, "testData", e1.Data, "Event 1 data is not correct")
	assert.Equal(t, "testData", e2.Data, "Event 2 data is not correct")

}
