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
	return EventEmitter{
		master:             make(map[string]map[string]Listener),
		maxListeners:       maxListeners,
		internalEventNames: []string{"newListener"},
	}
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
