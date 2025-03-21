package main

import "testing"

func Test_CheckWinner(t *testing.T) {
	testCases := map[string]struct {
		boardPath string
		winner    Cell // CellEmpty if no winner
		row       int  // 0 if no winner
		col       int
	}{
		"black wins, horizontal": {
			boardPath: "boards/1.csv",
			winner:    CellBlack,
			row:       17,
			col:       1,
		},
		"black wins, vertical": {
			boardPath: "boards/2.csv",
			winner:    CellBlack,
			row:       6,
			col:       1,
		},
		"black wins, \\": {
			boardPath: "boards/3.csv",
			winner:    CellBlack,
			row:       1,
			col:       1,
		},
		"no winner": {
			boardPath: "boards/4.csv",
		},
		"no winner, strike overflow": {
			boardPath: "boards/5.csv",
		},
		"black wins, /": {
			boardPath: "boards/6.csv",
			winner:    CellBlack,
			row:       8,
			col:       8,
		},
		"no winner, lots of overflow": {
			boardPath: "boards/7.csv",
		},
		"black winner": {
			boardPath: "boards/8.csv",
			winner:    CellBlack,
			row:       7,
			col:       2,
		},
		"white winner, \\": {
			boardPath: "boards/9.csv",
			winner:    CellWhite,
			row:       5,
			col:       5,
		},
		"black winner, /": {
			boardPath: "boards/10.csv",
			winner:    CellBlack,
			row:       19,
			col:       15,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			board, err := ReadRenjuBoard(tc.boardPath)
			if err != nil {
				t.Fatalf("failed to read board from %s: %v", tc.boardPath, err)
			}

			winner, row, col := CheckRenjuWinner(board)
			if winner != tc.winner {
				t.Errorf("expected winner %d, got %d", tc.winner, winner)
			}
			if col != tc.col {
				t.Errorf("expected winning column %d, got %d", tc.col, col)
			}
			if row != tc.row {
				t.Errorf("expected winning row %d, got %d", tc.row, row)
			}
		})
	}
}
