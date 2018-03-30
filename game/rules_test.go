package game

import "testing"

func TestHorizontalCheck(t *testing.T) {
	testGame := newGame()
	testGame.grid = Grid{
		top: [gridSize]rune{o, o, o},
		mid: [gridSize]rune{},
		bot: [gridSize]rune{},
	}

	expect := testGame.horizontalCheck()

	if expect == true {
		t.Logf("The winner is: ")
	} else {
		t.Errorf("FAIL horizontalCheck")
	}

}

func TestHorizontalCheckFail(t *testing.T) {
	testGame := newGame()
	testGame.grid = Grid{
		top: [gridSize]rune{o},
		mid: [gridSize]rune{x},
		bot: [gridSize]rune{x, x},
	}

	expect := testGame.horizontalCheck()

	if expect == false {
		t.Logf("No Winners today.")
	} else {
		t.Errorf("")
	}

}

func TestVerticalCheck(t *testing.T) {
	testGame := newGame()
	testGame.grid = Grid{
		top: [gridSize]rune{o, x, o},
		mid: [gridSize]rune{x, x, o},
		bot: [gridSize]rune{o, x, x},
	}

	expect := testGame.verticalCheck()

	if expect == true {
		t.Logf("The winner is: ")
	} else {
		t.Errorf("")
	}

}

func TestVerticalCheckFail(t *testing.T) {
	testGame := newGame()
	testGame.grid = Grid{
		top: [gridSize]rune{o, x},
		mid: [gridSize]rune{x},
		bot: [gridSize]rune{x, x, o},
	}

	expect := testGame.horizontalCheck()

	if expect == false {
		t.Logf("No winners today.")
	} else {
		t.Errorf("")
	}

}

func TestDiagonalCheck(t *testing.T) {
	testGame := newGame()
	testGame.grid = Grid{
		top: [gridSize]rune{x, x, o},
		mid: [gridSize]rune{o, x, o},
		bot: [gridSize]rune{o, x, x},
	}

	expect := testGame.verticalCheck()

	if expect == true {
		t.Logf("The winner is:")
	} else {
		t.Errorf("")
	}

}

func TestDiagonalCheckFail(t *testing.T) {
	testGame := newGame()
	testGame.grid = Grid{
		top: [gridSize]rune{o, x},
		mid: [gridSize]rune{x, x, o},
		bot: [gridSize]rune{x, x, o},
	}

	expect := testGame.horizontalCheck()

	if expect == false {
		t.Logf("No winners today.")
	} else {
		t.Errorf("")
	}

}
