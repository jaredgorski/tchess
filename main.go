package main

import (
	// "fmt"

	"github.com/jaredgorski/tchess/internal/board"
	"github.com/jaredgorski/tchess/internal/pieces"
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
	squares := [64]board.Square{}

	inc := 0;
	for i := 0; i < 64; i++ {
		piece := "BK"

		if inc == 1 {
			piece = "WP"
		} else if inc == 2 {
			piece = "WK"
		} else if inc == 3 {
			piece = "BB"
		} else if inc == 4 {
			piece = "_"
		}

		squares[i] = board.Square{
			IsHighlighted: false,
			Piece: pieces.Pieces[piece],
		}

		inc++

		if inc > 4 {
			inc = 0
		}
	}

	b := board.Board{
		IsWhiteSide: false,
		IsLarge: true,
		UseIcons: true,
		LastSquare: 1,
		Squares: squares,
	}

	board.DrawBoard(b)
}
