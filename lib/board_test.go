package lib

import "testing"

func TestNewBoard(t *testing.T) {
	board := createDefaultBoard()

	for i, c := range board.cells {
		for j, r := range c {
			if val, _ := r.Value(); val != board.cells[i][j].value {
				t.Errorf("invalid cell value")
			}
		}
	}
}

func TestOpenCell(t *testing.T) {
	board := createDefaultBoard()

	board.OpenCell(0, 0)
}

func createDefaultBoard() Board {
	values := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	board, _ := NewBoard(3, func(c, r int) Cell {
		cell, _ := NewCell(values[c][r])
		return cell
	})

	return board
}
