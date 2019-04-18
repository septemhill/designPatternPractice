package composite

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

