package observer

import "fmt"

type Event struct {
	Data int64
}

type Observer interface {
	OnNotify(Event)
}

type Notifier interface {
	Register(Observer)
	Unregister(Observer)
	Notify(Event)
}

type EventObserver struct {
}

type EventNotifier struct {
	m map[Observer]struct{}
}

func (o EventObserver) OnNotify(e Event) {
	fmt.Println(e)
}

func (n EventNotifier) Register(o Observer) {
	n.m[o] = struct{}{}
}

func (n EventNotifier) Unregister(o Observer) {
	delete(n.m, o)
}

func (n EventNotifier) Notify(e Event) {
	for ob := range n.m {
		ob.OnNotify(e)
	}
}
