
# [Go] Mini Event Bus Library

This is a minimalistic event bus library implemented in Go. It supports basic subscribe, unsubscribe, and publish operations.

### Features
 - Thread-safe operations
 - Unique ID generation for subscriptions

### Instalation
Make sure that Go is installed on your computer. Type the following command in your terminal:

```shell
go get github.com/yuksbg/meb
```

### Usage

### Creating a new Event Bus
To create a new instance of the event bus, use the function NewEventBus().
```go
eventBus := meb.NewEventBus()
```

### Subscribing to an Event

To subscribe to an event, use the Subscribe method, it takes in two parameters - an eventType of string type and an observer which is a function accepting an Event.

```go
subscriptionID := eventBus.Subscribe("event1", func(e meb.Event) {
	fmt.Println("Received data: ", e.Data)
})
```
This will return a subscriptionID of the subscriber which can be used to unsubscribe later.

### Unsubscribing from an Event
To unsubscribe a subscriber from an event, use the Unsubscribe method with the eventType and the subscriptionID.

```go
eventBus.Unsubscribe("event1", subscriptionID)
```

### Publishing an Event
To publish an Event to all the subscribed observers, we use Publish function of Event Bus which requires an eventType and the data.

```go
eventBus.Publish("event1", "Test Data")
```

### Notes
This library is built around the core concept of an EventBus. 
An EventBus holds a mapping of event type names to handlers representing subscriptions. 
These handlers are functions that get called with any published data when the associated event type is published to. 
The thread-safe operations are ensured by utilizing Go's Mutex lock and unlock before accessing the subscriptions map.
