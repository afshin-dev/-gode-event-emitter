package gode

// ListenerFunc a function that can take any number of
// optional interface{} type inputs and never returns any output/value
// this type can implement with  generic but it is more nodejs like at this manner
type ListenerFunc func(args ...interface{})

// Listener is a type that  for registering a
// ful informational event handler by `master object`
type Listener struct {
	fn   ListenerFunc
	once bool
	id   string
}
