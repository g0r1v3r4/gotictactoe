package game

// TicTacToeRules is an interface representing the necessary methods to implement a tic-tac-toe game
type TicTacToeRules interface {
	horizontalCheck()
	verticalCheck()
	diagonalCheck()
}

func (game TicTacToeGame) horizontalCheck() bool {
	switch {
	case game.grid.top[topRow] == game.grid.top[midRow] && game.grid.top[midRow] == game.grid.top[botRow]:
		return true
	case game.grid.mid[topRow] == game.grid.mid[midRow] && game.grid.mid[midRow] == game.grid.mid[botRow]:
		return true
	case game.grid.bot[topRow] == game.grid.bot[midRow] && game.grid.bot[midRow] == game.grid.bot[botRow]:
		return true
	default:
		return false
	}
}

func (game TicTacToeGame) verticalCheck() bool {
	for index := range game.grid.top {
		if game.grid.top[index] == game.grid.mid[index] && game.grid.mid[index] == game.grid.bot[index] {
			return true
		}
	}
	return false
}

func (game TicTacToeGame) diagonalCheck() bool {
	switch {
	case game.grid.top[topRow] == game.grid.mid[midRow] && game.grid.mid[midRow] == game.grid.bot[botRow]:
		return true
	case game.grid.top[botRow] == game.grid.mid[midRow] && game.grid.mid[midRow] == game.grid.bot[topRow]:
		return true
	default:
		return false
	}
}
