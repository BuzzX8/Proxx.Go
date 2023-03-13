package lib

import "errors"

type Cell struct {
	value  int
	isHole bool
	isOpen bool
}

func NewCell(value int) (Cell, error) {
	if value < 0 {
		return Cell{}, errors.New("invalid cell value")
	}
	return Cell{value: value}, nil
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

func (cell *Cell) Value() (int, error) {
	if cell.isHole {
		return 0, errors.New("this is hole")
	}

	return cell.value, nil
}
