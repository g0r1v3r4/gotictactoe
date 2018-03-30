package game

import "testing"

func TestNewPlayer(t *testing.T) {

	expect := newPlayer(x)

	if expect.symbol == x {
		t.Logf("Test player created successfully %v", expect)
	} else {
		t.Errorf("Failed to create player successfully.")
	}

}
