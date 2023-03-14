package lib

import "errors"

type CellGenerator func(column, row int) Cell

type GameState byte

const (
	InProgress GameState = 0
	Lost       GameState = 1
	Won        GameState = 2
)

type Board struct {
	cells []Cell
	size  int
}

func NewBoard(size int, generator CellGenerator) (Board, error) {
	if size <= 2 {
		return Board{}, errors.New("invalid board size")
	}

	cells := make([]Cell, size*size)

	for column := 0; column < size; column++ {
		for row := 0; row < size; row++ {
			cells[getIndex(column, row, size)] = generator(column, row)
		}
	}

	return Board{cells: cells, size: size}, nil
}

func (board *Board) GetCell(column, row int) (*Cell, error) {
	if column > board.size || row > board.size {
		return nil, errors.New("index out of bounds")
	}

	return &board.cells[getIndex(column, row, board.size)], nil
}

func (board *Board) GameState() GameState {
	for _, cell := range board.cells {
		if cell.isHole && cell.isOpen {
			return Lost
		}
	}

	for _, cell := range board.cells {
		if !cell.isHole && !cell.isOpen {
			return InProgress
		}
	}

	return Won
}

func (board *Board) OpenCell(column, row int) error {
	if column >= board.size || row >= board.size {
		return errors.New("invalid arguments")
	}

	cell := &board.cells[getIndex(column, row, board.size)]
	cell.isOpen = true
	adjCells := board.getAdjectentCells(column, row)

	for _, adjCell := range adjCells {
		if adjCell.isHole {
			return nil
		}
	}

	for _, adjCell := range adjCells {
		adjCell.isOpen = true
	}

	return nil
}

func (board *Board) getAdjectentCells(column, row int) []*Cell {
	var adjCells []*Cell

	for cOff := -1; cOff <= 1; cOff++ {
		for rOff := -1; rOff <= 1; rOff++ {
			cIndx := column + cOff
			rIndx := row + rOff

			if cIndx < 0 || cIndx >= board.size || rIndx < 0 || rIndx >= board.size {
				continue
			}

			if cIndx == column && rIndx == row {
				continue
			}

			adjCells = append(adjCells, &board.cells[getIndex(column, row, board.size)])
		}
	}

	return adjCells
}

func getIndex(column, row, size int) int {
	return size*column + row
}

func (board *Board) Size() int {
	return board.size
}
