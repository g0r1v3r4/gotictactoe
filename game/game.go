package game

import (
	"fmt"
)

// gridSize standard tic-tac-toe sized grid
const gridSize int = 3

const (
	topRow = iota
	midRow
	botRow
)

const (
	leftColumn = iota
	midColumn
	rightColumn
)

// Grid represents the 2D array for the tic-tac-toe game and uses GridSize to create it
type Grid struct {
	top, mid, bot [gridSize]rune
}

// Coordinate value representation of the tic-tac-toe grid as a coordinate pair
type Coordinate struct {
	column int
	row    int
}

// Game defines the most basic methods required for games
type Game interface {
	Loop()
}

// TicTacToeGame declares the instance of a game and related functions
type TicTacToeGame struct {
	grid             Grid
	playerTurn       bool
	player1, player2 Player
}

func newGame() *TicTacToeGame {
	newGame := &TicTacToeGame{}
	newGame.playerTurn = true // player1: true, player2: false
	newGame.player1 = newPlayer(o)
	newGame.player2 = newPlayer(x)
	newGame.grid = Grid{
		top: [gridSize]rune{},
		mid: [gridSize]rune{},
		bot: [gridSize]rune{},
	}

	return newGame
}

func (game TicTacToeGame) isWinner() (bool, error) {
	// var (
	// 	diagonal, vertical, horizontal bool
	// )

	// diagonal = game.diagonalCheck()
	// horizontal = game.horizontalCheck()
	// vertical = game.verticalCheck()
	// TODO if who returns zero-value for rune panic()
	// There should not be an instance where rune has zero-value

	// if diagonal || vertical || horizontal {
	// 	// A win event has occurred; time to see who won
	// 	fmt.Printf("Winner is: %q", who)

	// 	if game.player1.symbol == who {
	// 		return game.player1, nil
	// 	} else if game.player2.symbol == who {
	// 		// this check can be eliminated, but attempting to be explicit
	// 		return game.player2, nil
	// 	}
	// }

	return true, nil
}

// Loop represents the entire game loop for use exeternally
func Loop() {

	var entry Coordinate

	fmt.Println("Hello, Welcome to Tic Tac Toe")
	newGame := newGame()

	RenderGrid(newGame.grid)

	fmt.Println("Player1 will be O.\nPlayer2 will be X.\n" +
		"Use `row, column`. Must be integers from 0 - 2 separated with a comma.\n" +
		"Enter grid coordinates for your move:")

	if newGame.playerTurn == true {
		fmt.Println("Player1 it is  your turn:")

		inputs, err := fmt.Scanf("%d, %d", &entry.row, &entry.column)
		if inputs != 2 || err != nil {
			fmt.Println("What is input: ", inputs, entry.row, entry.column)
			if entry.row >= 0 && entry.row < gridSize {
				fmt.Errorf("Out of bounds row")
			}
			if entry.column >= 0 && entry.column < gridSize {
				fmt.Errorf("Out of bounds column")
			}
		}

		symbol := newGame.player1.symbol

		if newGame.isCellBlank(entry) {
			switch entry.row {
			case topRow:
				newGame.grid.top[entry.column] = symbol
			case midRow:
				newGame.grid.top[entry.column] = symbol
			case botRow:
				newGame.grid.top[entry.column] = symbol
			default:
				fmt.Errorf("Something went wrong.")
			}
		}
		//newGame.makeMove(entry, newGame.player1)

	} else {
		fmt.Println("Player2 it is  your turn:")
	}

	// for !g.isWinner(g.player1) {
	// 	start()
	// }
}

// RenderGrid prints to standard out the tic-tac-toe grid
func RenderGrid(g Grid) {
	fmt.Printf("| %s | %s | %s |\n", string(g.top[leftColumn]), string(g.top[midColumn]), string(g.top[rightColumn]))
	fmt.Printf("|---|---|---|\n")
	fmt.Printf("| %s | %s | %s |\n", string(g.mid[leftColumn]), string(g.mid[midColumn]), string(g.mid[rightColumn]))
	fmt.Printf("|---|---|---|\n")
	fmt.Printf("| %s | %s | %s |\n", string(g.bot[leftColumn]), string(g.bot[midColumn]), string(g.bot[rightColumn]))
}

// isCellBlank determines if a cell in the grid is empty
func (game TicTacToeGame) isCellBlank(coordinate Coordinate) bool {
	switch coordinate.row {
	case topRow:
		return game.grid.top[coordinate.column] == 0
	case midRow:
		return game.grid.mid[coordinate.column] == 0
	case botRow:
		return game.grid.bot[coordinate.column] == 0
	default:
		return false
	}
}
