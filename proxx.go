package proxx

import "errors"

type Cell struct {
	value  uint
	isHole bool
	isOpen bool
}

func NewCell(value uint) Cell {

	return Cell{value: value}
}

func NewHoleCell() Cell {
	return Cell{isHole: true}
}

func (cell *Cell) IsHole() bool {
	return cell.isHole
}

func (cell *Cell) IsOpen() bool {
	return cell.isOpen
}

func (cell *Cell) Value() (uint, error) {
	if cell.isHole {
		return 0, errors.New("this is hole")
	}

	return cell.value, nil
}

type CellGenerator func(column, row uint) Cell

type GameState byte

const (
	InProgress GameState = 0
	Lost       GameState = 1
	Won        GameState = 2
)

type Board struct {
	cells [][]Cell
	size  uint
}

func NewBoard(size uint, generator CellGenerator) Board {
	cells := make([][]Cell, size)

	for column := uint(0); column < size; column++ {
		cells[column] = make([]Cell, size)
		for row := uint(0); row < size; row++ {
			cells[column][row] = generator(column, row)
		}
	}

	return Board{cells: cells, size: size}
}

func (board *Board) GetCell(column, row uint) (*Cell, error) {
	if column > board.size || row > board.size {
		return nil, errors.New("index out of bounds")
	}

	return &board.cells[column][row], nil
}

func (board *Board) GameState() GameState {
	panic("not implemented")
}

func (board *Board) OpenCell(column, row uint) error {
	panic("not implemented")
}
