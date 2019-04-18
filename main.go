package main

import (
	"dp/behavioral/state"
)

func main() {
	start := state.StartState{}
	game := state.GameContext{
		Next: &start,
	}

	for game.Next.ExecuteState(&game) {
	}
}
