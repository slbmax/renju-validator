package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please, provide a file path to the board configuration")
		return
	}

	board, err := ReadRenjuBoard(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	winner, row, col := CheckRenjuWinner(board)
	if winner == CellEmpty {
		fmt.Println("No winner")
	} else {
		fmt.Printf("Winner: %d at (%d,%d)\n", winner, row, col)
	}
}

func ReadRenjuBoard(filePath string) ([BoardSize][BoardSize]Cell, error) {
	var board [BoardSize][BoardSize]Cell

	file, err := os.Open(filePath)
	if err != nil {
		return board, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read and parse each line
	for x := 0; x < BoardSize; x++ {
		if !scanner.Scan() {
			return board, fmt.Errorf("file contains fewer than %d lines", BoardSize)
		}

		line := scanner.Text()
		values := strings.Split(line, ",")

		if len(values) != BoardSize {
			return board, fmt.Errorf("line %d contains %d values, expected %d", x+1, len(values), BoardSize)
		}

		for y := 0; y < BoardSize; y++ {
			// Trim any whitespace
			valueStr := strings.TrimSpace(values[y])

			// Parse the cell value
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return board, fmt.Errorf("invalid value at position (%d,%d): %s", x+1, y+1, valueStr)
			}
			if value < 0 || value > 2 {
				return board, fmt.Errorf("invalid cell value at position (%d,%d): %d (must be 0, 1, or 2)", x+1, y+1, value)
			}

			board[x][y] = Cell(value)
		}
	}

	if scanner.Scan() {
		return board, fmt.Errorf("file contains more than %d lines", BoardSize)
	}
	if err := scanner.Err(); err != nil {
		return board, fmt.Errorf("error reading file: %w", err)
	}

	return board, nil
}

func CheckRenjuWinner(board [BoardSize][BoardSize]Cell) (winner Cell, col, row int) {
	for x := range BoardSize {
		for y := range BoardSize {
			cell := board[x][y]
			if cell == CellEmpty {
				continue
			}

			// checking for a win in all possible directions
			for _, direction := range Directions {
				if direction.WinImpossible(x, y) {
					continue
				}

				if !isWinStrike(board, x, y, direction) {
					continue
				}

				leftMostX, leftMostY := direction.LeftMost(x, y)

				// returning coordinates, not indexes
				return cell, leftMostX + 1, leftMostY + 1
			}
		}
	}

	return CellEmpty, 0, 0
}

// isWinStrike checks if there is a win strike starting from the cell (x, y) in the direction (dx, dy).
// It assumes that all the cells in the strike are within the board
func isWinStrike(board [BoardSize][BoardSize]Cell, x, y int, direction Direction) bool {
	cell := board[x][y]
	dx, dy := direction.DxDy()

	for i := 1; i < WinStrike; i++ {
		nextX, nextY := x+i*dx, y+i*dy
		if board[nextX][nextY] != cell {
			return false
		}
	}

	// checking for the strike overflows
	prevX, prevY := x-dx, y-dy
	if WithinTheBoard(prevX, prevY) && board[prevX][prevY] == cell {
		return false
	}

	nextX, nextY := x+WinStrike*dx, y+WinStrike*dy
	if WithinTheBoard(nextX, nextY) && board[nextX][nextY] == cell {
		return false
	}

	return true
}
