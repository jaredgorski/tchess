package main

import (
	// "fmt"

	"github.com/jaredgorski/tchess/internal/board"
)

var (
	successExit	= 0
	errorExit	= 1
)

type silentfail struct {}

func (silentfail) Error() string {
	return ""
}

func main() {
	b := board.Board{
		IsWhiteSide: true,
		IsLarge: true,
		UseIcons: true,
		LastSquare: 1,
	}

	b.ResetBoard()
	b.DrawBoard()

	b.MovePiece("Pd2d3")
}
