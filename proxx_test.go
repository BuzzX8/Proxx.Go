package proxx

import "testing"

func TestNewBoard(t *testing.T) {
	values := [][]uint{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	board := NewBoard(3, func(c, r uint) Cell {
		return NewCell(values[c][r])
	})

	for i, c := range board.cells {
		for j, r := range c {
			if val, _ := r.Value(); val != values[i][j] {
				t.Errorf("invalid cell value")
			}
		}
	}
}
