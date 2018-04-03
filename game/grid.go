package game

import (
	"errors"
	"fmt"
)

// gridSize standard tic-tac-toe sized grid (demonstrates Go constants)
const gridSize int = 3

// *Row are navegational indexes within Grid
const (
	topRow = iota
	midRow
	botRow
)

// *Column are navegational indexes within Grid (demonstrates Go enums)
const (
	leftColumn = iota
	midColumn
	rightColumn
)

// Coordinate value representation of the tic-tac-toe grid as a coordinate pair (demonstrates exported type)
type Coordinate struct {
	column int
	row    int
}

// Grid represents the 2D array for the tic-tac-toe game (demonstrates struc definitions with Go array)
type Grid struct {
	top, mid, bot [gridSize]rune
}

// Unexported interface represents the functionality of type Grid (demonstrate unexported type interface)
type grid interface {
	RenderGrid(g *Grid)
	insertAtCoordinate(c Coordinate, symbol rune) (bool, error)
	getCellValue(c Coordinate) (rune, error)
	isCellBlank(c Coordinate) bool
}

// constructor for type Grid
func newGrid() *Grid {
	return &Grid{
		top: [gridSize]rune{},
		mid: [gridSize]rune{},
		bot: [gridSize]rune{},
	}
}

// RenderGrid prints to standard out the tic-tac-toe grid (demonstrates exported function)
func RenderGrid(g *Grid) {
	fmt.Printf("|-00-|-01-|-02-|\n")
	fmt.Printf("| %s | %s | %s |\n", string(g.top[leftColumn]), string(g.top[midColumn]), string(g.top[rightColumn]))
	fmt.Printf("|-10-|-11-|-12-|\n")
	fmt.Printf("| %s | %s | %s |\n", string(g.mid[leftColumn]), string(g.mid[midColumn]), string(g.mid[rightColumn]))
	fmt.Printf("|-20-|-21-|-22-|\n")
	fmt.Printf("| %s | %s | %s |\n", string(g.bot[leftColumn]), string(g.bot[midColumn]), string(g.bot[rightColumn]))
	fmt.Printf("|-EG-|-EG-|-EG-|\n")
}

// isCellBlank determines if a cell in the grid is empty (demonstrates type method with receiver value)
func (grid Grid) isCellBlank(c Coordinate) bool {
	switch c.row {
	case topRow:
		return grid.top[c.column] == 0
	case midRow:
		return grid.mid[c.column] == 0
	case botRow:
		return grid.bot[c.column] == 0
	default:
		return false
	}
}

// insertAtCoordinate inserts the symbol(rune) into the valid c(Coordinate) within grid(Grid)
// (demonstrates error handling, input validation, and multi-value return, adress references)
func (grid *Grid) insertAtCoordinate(c Coordinate, symbol rune) (bool, error) {
	if c.row >= gridSize && c.row < 0 || c.column >= gridSize && c.column < 0 {
		return false, errors.New("row or column in coordinate specified is out of bounds for array[gridSize]")
	}

	switch c.row {
	case topRow:
		grid.top[c.column] = symbol
	case midRow:
		grid.mid[c.column] = symbol
	case botRow:
		grid.bot[c.column] = symbol
	default:
		return false, errors.New("could not assign symbol to grid")
	}

	return true, nil
}

func (grid Grid) getCellValue(c Coordinate) (rune, error) {
	if c.row >= gridSize && c.row < 0 || c.column >= gridSize && c.column < 0 {
		return 0, errors.New("row or column in coordinate specified is out of bounds for array[gridSize]")
	}

	switch c.row {
	case topRow:
		return grid.top[c.column], nil
	case midRow:
		return grid.mid[c.column], nil
	case botRow:
		return grid.bot[c.column], nil
	default:
		return 0, errors.New("something went wrong")
	}
}
