package main

import (
	"fmt"
	"proxx/proxx"
)

func main() {
	board := generateBoard(3, 3)

	for board.GameState() == proxx.InProgress {
		panic("not implemented")
	}

	switch board.GameState() {
	case proxx.Lost:
		fmt.Println("You lost")
	case proxx.Won:
		fmt.Println("You won")
	}
}

func generateBoard(size, holeCount uint) proxx.Board {
	return proxx.NewBoard(size, func(column, row uint) proxx.Cell {
		panic("not implemented")
	})
}
