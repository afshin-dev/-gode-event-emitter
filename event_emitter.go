package gode

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math"
)

const maxListeners int = math.MaxInt

// EventEmitter is main type that you need to use for
// using event emitter functionality
type EventEmitter struct {
	// master is a map that every key mapped to a map that keys mapped to Listener
	master             map[string]map[string]Listener
	maxListeners       int
	internalEventNames []string
}

// New return a new instance of EventEmitter
func New() EventEmitter {
	eventEmitter := EventEmitter{
		master:       make(map[string]map[string]Listener),
		maxListeners: maxListeners,
		// TODO: implement these internal event into right place
		internalEventNames: []string{"newListener", "removeListener"},
	}

	// register default and internal events that EventEmitter at some point emit
	// for example whenever any new event listener added to master
	// EventEmitter emit a `newListener` for everyone listen to it
	for _, internalEvent := range eventEmitter.internalEventNames {
		eventEmitter.master[internalEvent] = make(map[string]Listener)
	}

	return eventEmitter
}

// Listeners return a slice of ListenerFunc and if this listeners call
// manually by consumers, by definition or ideally must not cause any side effect in
// internal of EventEmitter
func (ee *EventEmitter) Listeners(eventName string) []ListenerFunc {
	var listenerFuncs []ListenerFunc

	for _, listener := range ee.master[eventName] {
		listenerFuncs = append(listenerFuncs, listener.fn)
	}

	return listenerFuncs
}

// ListenerCount return len of listeners map for provided eventName
func (ee *EventEmitter) ListenerCount(eventName string) int {
	listeners, ok := ee.master[eventName]

	if !ok {
		return 0
	}
	return len(listeners)
}

// EventNames return any event name that EventEmitter has registered listener for it
func (ee *EventEmitter) EventNames() []string {
	var events []string
	for event, listenerMap := range ee.master {
		if len(listenerMap) > 0 {
			events = append(events, event)
		}
	}
	return events
}

// AddEventListener add a listener to underlying master events object
func (ee *EventEmitter) AddEventListener(eventName string, listenerFunc ListenerFunc) (listenerId string) {
	listenerId = uuid.NewString()

	listener := Listener{
		fn:   listenerFunc,
		once: false,
		id:   listenerId,
	}

	// if is the first time that a event going to create
	if _, ok := ee.master[eventName]; !ok {
		ee.master[eventName] = make(map[string]Listener)
		ee.master[eventName][listenerId] = listener

		return
	}

	// check for max_listener
	ee.checkMaxListeners(eventName)

	ee.master[eventName][listenerId] = listener
	return
}

func (ee *EventEmitter) RemoveEventListener(eventName string, listenerId string) {
	if _, ok := ee.master[eventName]; !ok {
		return
	}

	delete(ee.master[eventName], listenerId)

	// Revision: if not deleting a empty listener map cause any performance or memory
	// problem `here` after deleting event listener must delete empty map
}

// Emit call all listener as a goroutine for provided event name
// all listener must be pure and not create side effect
// listener better to not panic
// and when accessing global state you should access synchronously
// with some thing like lock
func (ee *EventEmitter) Emit(eventName string, args ...interface{}) {
	for _, l := range ee.master[eventName] {
		go l.fn(args)
	}
}

func (ee *EventEmitter) checkMaxListeners(eventName string) {
	if len(ee.master[eventName]) >= ee.maxListeners {
		panic(errors.New(fmt.Sprintf("max listener exceeded : max_listener number is : %d",
			ee.maxListeners)))
	}
}
func (ee *EventEmitter) GetMaxListeners() int {
	return ee.maxListeners
}
func (ee *EventEmitter) SetMaxListeners(newMaxListeners int) {
	// if except internal event already registered another event
	// SetMaxListeners must not be set a number in EventEmitter instance
	if len(ee.master) > len(ee.internalEventNames) {
		panic("set listeners must be call at initial phase")
	}
	ee.maxListeners = newMaxListeners
}
