package memento

import "testing"

func TestApplication(t *testing.T) {
	app := NewApplication()

	s1 := State{
		HealthPoint: 100,
		ManaPoint:   500,
	}

	t.Run("Application - SetState/GetState", func(t1 *testing.T) {
	})

	t.Run("Application - Backup", func(t1 *testing.T) {
	})

	app.SetState(s1)
}
func TestMemento(t *testing.T) {
	app := NewApplication()
	recovery := NewRecovery()

	app.SetState(State{
		HealthPoint: 100,
		ManaPoint:   50,
	})

	bak := app.Backup()

	recovery.AddBackup(bak)

	app.SetState(State{
		HealthPoint: 454,
		ManaPoint:   567,
	})

	k := recovery.Retrieve(bak.Timestamp)

	app.SetState(k.State)

	g := app.GetState()

	if g.HealthPoint != 100 || g.ManaPoint != 50 {
		t.Error("NOOOO")
	}
}
