package main

import (
	"fmt"
	"math/rand"
	"proxx/proxx"
	"strconv"
)

func main() {
	board := generateBoard(3, 3)

	for board.GameState() == proxx.InProgress {
		renderBoard(&board)
		column, row := getUserInput()
		board.OpenCell(column, row)
	}

	switch board.GameState() {
	case proxx.Lost:
		fmt.Println("You lost")
	case proxx.Won:
		fmt.Println("You won")
	}
}

func getUserInput() (column, row uint) {
	var colStr, rowStr string

	fmt.Print("Enter column index: ")
	fmt.Scanln(&colStr)

	c, _ := strconv.ParseUint(colStr, 10, 64)

	fmt.Print("Enter row index: ")
	fmt.Scanln(&rowStr)

	r, _ := strconv.ParseUint(rowStr, 10, 64)

	column, row = uint(c), uint(r)
	return
}

func generateBoard(size, holeCount uint) proxx.Board {
	holesPosition := make([]struct{ column, row uint }, holeCount)

	for i := 0; i < int(holeCount); i++ {

	}

	return proxx.NewBoard(size, func(column, row uint) proxx.Cell {
		for _, hole := range holesPosition {
			if hole.column == column && hole.row == row {
				return proxx.NewHoleCell()
			}
		}
		return proxx.NewCell(uint(rand.Intn(4)))
	})
}

func renderBoard(board *proxx.Board) {
	for row := uint(0); row < board.Size(); row++ {
		for column := uint(0); column < board.Size(); column++ {
			cell, _ := board.GetCell(column, row)
			fmt.Print("|")
			renderCell(cell)
		}
		fmt.Print("|")
		fmt.Println()
	}
}

func renderCell(cell *proxx.Cell) {
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
