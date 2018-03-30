package game

import "testing"

func TestNewGame(t *testing.T) {
	expect := newGame()

	if expect != nil {
		t.Logf("Test grid created successfully.")
	} else {
		t.Errorf("Failed to create grid.")
	}

	if expect.player1.symbol == o && expect.player2.symbol == x {
		t.Logf("Test player created successfully")
	} else {
		t.Errorf("Failed to create player successfully.")
	}
}
