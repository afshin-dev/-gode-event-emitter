package gode

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestEventEmitter_AddEventListener_Must_Return_A_ListenerId(t *testing.T) {
	ee := New()
	listenerId := ee.AddEventListener("some_event", func(args ...interface{}) {

	})

	if _, err := uuid.Parse(listenerId); err != nil {
		t.Error("listener must be return a valid string uuid")
	}
}

func TestEventEmitter_GetMaxListeners_Must_Return_Constant_maxListeners(t *testing.T) {
	ee := New()
	if ee.GetMaxListeners() != maxListeners {
		t.Error("by default EventEmitter maxListeners set to constant maxListeners")
	}
}

func TestEventEmitter_SetMaxListeners_Must_Change_At_Phase_Of_Initial(t *testing.T) {
	ee := New()
	newMaxListeners := 1000

	ee.SetMaxListeners(newMaxListeners)

	if ee.GetMaxListeners() != newMaxListeners {
		t.Error(fmt.Printf("newMaxListeners must be %d", newMaxListeners))
	}
}

func TestEventEmitter_EventNames_Registered_Event_Must_Included_In_Returns_Slice(t *testing.T) {
	ee := New()
	exampleEventName := "a"
	ee.AddEventListener(exampleEventName, ListenerFunc(func(args ...interface{}) {

	}))

	names := ee.EventNames()

	for _, e := range names {
		if e == exampleEventName {
			return
		}
	}

	t.Error(fmt.Sprintf("%s must be in event  names slice", exampleEventName))
}
