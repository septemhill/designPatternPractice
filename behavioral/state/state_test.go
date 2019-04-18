package state

import "testing"

func TestGameState(t *testing.T) {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}

	for game.Next.executeState(&game) {
	}
}
