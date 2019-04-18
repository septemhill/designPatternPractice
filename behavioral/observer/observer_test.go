package observer

import "testing"

func TestObserver(t *testing.T) {
	var ob1, ob2 EventObserver
	n := EventNotifier{
		m: make(map[Observer]struct{}),
	}

	n.Register(ob1)
	n.Register(ob2)

	n.Notify(Event{Data: 100})
}
