package memento

import (
	"time"
)

type State struct {
	HealthPoint int
	ManaPoint   int
}

type Backup struct {
	State     State
	Timestamp int64
}

func NewBackup(state State) *Backup {
	b := &Backup{
		State:     state,
		Timestamp: time.Now().UnixNano(),
	}

	return b
}

type Application struct {
	state State
}

func (a *Application) Backup() *Backup {
	return NewBackup(a.state)
}

func (a *Application) Recover(bak Backup) {
	a.state = bak.State
}

func (a *Application) SetState(state State) {
	a.state = state
}

func (a *Application) GetState() State {
	return a.state
}

func NewApplication() *Application {
	a := &Application{}

	return a
}

type Recovery struct {
	backups map[int64]*Backup
}

func NewRecovery() *Recovery {
	r := &Recovery{
		backups: make(map[int64]*Backup),
	}
	return r
}

func (r *Recovery) AddBackup(bak *Backup) {
	r.backups[bak.Timestamp] = bak
}

func (r *Recovery) Retrieve(time int64) *Backup {
	bak, ok := r.backups[time]

	if ok {
		delete(r.backups, time)
	}

	return bak
}

//type State struct {
//	Description string
//}
//
//type memento struct {
//	state State
//}
//
//type originator struct {
//	state State
//}
//
//func (o *originator) NewMemento() memento {
//	return memento{}
//}
//
//func (o *originator) ExtractAndStoreState(m memento) {
//
//}
//
//type careTaker struct {
//	mementoList []memento
//}
//
//func (c *careTaker) Add(m memento) {
//	c.mementoList = append(c.mementoList, m)
//}
//
//func (c *careTaker) Memento(i int) (memento, error) {
//	return memento{}, fmt.Errorf("not implemented yet")
//}
