package composite

import "testing"

func TestComposite(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmer.Swim()
	swimmer.Train()
}
