package proxx

import "errors"

type Cell struct {
	value  int
	isOpen bool
}

func NewCell(value int) (Cell, error) {
	if value < 0 {
		return Cell{}, errors.New("argument exception")
	}

	return Cell{value: value, isOpen: false}, nil
}

func NewHoleCell() Cell {
	return Cell{value: -1, isOpen: false}
}

func (cell *Cell) IsHole() bool {
	return cell.value <= 0
}

func (cell *Cell) IsOpen() bool {
	return cell.isOpen
}

func (cell *Cell) Value() (int, error) {
	if cell.value < 0 {
		return 0, errors.New("this is hole")
	}

	return cell.value, nil
}

type CellGenerator func(column, row uint) Cell

type Board struct {
	cells [][]Cell
	size  uint
}

func NewBoard(size uint, generator CellGenerator) Board {
	cells := make([][]Cell, size)

	for column := uint(0); column < size; column++ {
		for row := uint(0); row < size; row++ {
			cells[column][row] = generator(column, row)
		}
	}

	return Board{cells: cells, size: size}
}
