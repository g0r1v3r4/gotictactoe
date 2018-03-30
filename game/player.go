package game

const (
	x rune = 'X'
	o rune = 'O'
)

// Player definition
type Player struct {
	symbol      rune
	wins, loses int
}

// a constructor method for type Player
func newPlayer(_symbol rune) Player {
	player := Player{
		symbol: _symbol,
		wins:   0,
		loses:  0,
	}
	return player
}
