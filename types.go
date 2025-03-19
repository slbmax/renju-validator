package main

const (
	BoardSize = 19
	WinStrike = 5
)

// WithinTheBoard returns true if the given coordinates are within the board.
func WithinTheBoard(x, y int) bool {
	return 0 <= x && x < BoardSize && 0 <= y && y < BoardSize
}

// Cell is a type that represents a cell on the board.
type Cell int

const (
	CellEmpty Cell = iota
	CellBlack
	CellWhite
)

// Direction is a type that represents a direction on the board.
type Direction interface {
	// DxDy returns the change in x and y for each step in the direction.
	DxDy() (dx, dy int)
	// WinImpossible returns true if it is impossible to win in this direction from the given position.
	WinImpossible(x, y int) bool
	// LeftMost returns the coordinates of the leftmost cell in the direction
	// from the starting position for a win strike.
	LeftMost(x, y int) (leftMostX, leftMostY int)
}

// Directions is a list of all possible directions on the board.
var Directions = []Direction{
	Horizontal{},
	Vertical{},
	LeftDiagonal{},
	RightDiagonal{},
}

type Horizontal struct{}

func (Horizontal) DxDy() (dx, dy int) {
	return 0, 1
}

func (Horizontal) WinImpossible(x, y int) bool {
	return y > BoardSize-WinStrike
}

func (Horizontal) LeftMost(x, y int) (leftMostX, leftMostY int) {
	return x, y
}

type Vertical struct{}

func (Vertical) DxDy() (dx, dy int) {
	return 1, 0
}

func (Vertical) WinImpossible(x, y int) bool {
	return x > BoardSize-WinStrike
}

func (Vertical) LeftMost(x, y int) (leftMostX, leftMostY int) {
	return x, y
}

type LeftDiagonal struct{}

func (LeftDiagonal) DxDy() (dx, dy int) {
	return 1, 1
}

func (LeftDiagonal) WinImpossible(x, y int) bool {
	return x > BoardSize-WinStrike || y > BoardSize-WinStrike
}

func (LeftDiagonal) LeftMost(x, y int) (leftMostX, leftMostY int) {
	return x, y
}

type RightDiagonal struct{}

func (RightDiagonal) DxDy() (dx, dy int) {
	return 1, -1
}

func (RightDiagonal) WinImpossible(x, y int) bool {
	return x > BoardSize-WinStrike || y < WinStrike-1
}

func (RightDiagonal) LeftMost(x, y int) (leftMostX, leftMostY int) {
	return x + WinStrike - 1, y - WinStrike + 1
}
