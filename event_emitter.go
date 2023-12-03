package gode

import "math"

const max_listenres int = math.MaxInt

// EventEmitter is main type that you need to use for
// using event emitter functionality
type EventEmitter struct {
	// master is a map that every key mapped to a map that keys mapped to Listener
	master       map[string]map[string]Listener
	maxListeners int
}

// New return a new instance of EventEmitter
func New() EventEmitter {
	return EventEmitter{
		master:       make(map[string]map[string]Listener),
		maxListeners: max_listenres,
	}
}
