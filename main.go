package main

import (
	"fmt"
	"math/rand"
	"proxx/lib"
	"strconv"
)

func main() {
	board := generateBoard(3, 3)
	renderBoard(&board)

	for board.GameState() == lib.InProgress {
		column, row := getUserInput()
		board.OpenCell(column, row)
		renderBoard(&board)
	}

	switch board.GameState() {
	case lib.Lost:
		fmt.Println("You lost")
	case lib.Won:
		fmt.Println("You won")
	}
}

func getUserInput() (column, row int) {
	var colStr, rowStr string

	fmt.Print("Enter column index: ")
	fmt.Scanln(&colStr)

	c, _ := strconv.ParseInt(colStr, 10, 64)

	fmt.Print("Enter row index: ")
	fmt.Scanln(&rowStr)

	r, _ := strconv.ParseInt(rowStr, 10, 64)

	column, row = int(c), int(r)
	return
}

func generateBoard(size, holeCount int) lib.Board {
	holesPosition := make([]struct{ column, row int }, holeCount)

	for i := 0; i < int(holeCount); i++ {

	}

	board, _ := lib.NewBoard(size, func(column, row int) lib.Cell {
		for _, hole := range holesPosition {
			if hole.column == column && hole.row == row {
				return lib.NewHoleCell()
			}
		}
		cell, _ := lib.NewCell(rand.Intn(4))
		return cell
	})

	return board
}

func renderBoard(board *lib.Board) {
	for row := 0; row < board.Size(); row++ {
		for column := 0; column < board.Size(); column++ {
			cell, _ := board.GetCell(column, row)
			fmt.Print("|")
			renderCell(cell)
		}
		fmt.Print("|")
		fmt.Println()
	}
}

func renderCell(cell *lib.Cell) {
	switch {
	case !cell.IsOpen():
		fmt.Print(" ")
	case cell.IsHole():
		fmt.Print("X")
	default:
		v, _ := cell.Value()
		fmt.Print(v)
	}
}
