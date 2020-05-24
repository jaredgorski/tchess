package board

import (
	"fmt"
	"strings"

	"github.com/jaredgorski/tchess/internal/pieces"
)

var paddingSpace = string('\u00A0')

type Square struct {
	IsHighlighted	bool
	Piece			pieces.Piece
}

type Board struct {
	IsWhiteSide	bool
	IsLarge		bool
	UseIcons	bool
	LastSquare	int
	Squares		[64]Square
}

func center(s string, n int) string {
	div := n / 2

	return strings.Repeat(paddingSpace, div) + s + strings.Repeat(paddingSpace, div)
}

var files = [8]string{"A", "B", "C", "D", "E", "F", "G", "H"};
var ranks = [8]string{"1", "2", "3", "4", "5", "6", "7", "8"};

func DrawBoard(board Board) {
	flag := true
	width := 2

	if board.IsLarge {
		width = 5
	}

	fmt.Printf("\n");

	for i, rank := range ranks {
		flag = !flag

		if board.IsWhiteSide {
			rank = ranks[(len(ranks) - 1) - i]
		}

		if board.IsLarge {
			fmt.Print(center(paddingSpace, width))

			for i = 0; i < len(files); i++ {
				flag = !flag

				fmt.Printf("%v", DrawSquare(flag, paddingSpace, width))
			}

			fmt.Printf("\n");
		}

		fmt.Print(center(rank, width))

		for i = 0; i < len(files); i++ {
			flag = !flag

			square := board.Squares[i]
			icon := square.Piece.Name

			if board.UseIcons {
				icon = string(square.Piece.Icon)
			}

			fmt.Printf("%v", DrawSquare(flag, icon, width))
		}

		fmt.Printf("\n");

		if board.IsLarge {
			fmt.Print(center(paddingSpace, width))

			for i = 0; i < len(files); i++ {
				flag = !flag

				fmt.Printf("%v", DrawSquare(flag, paddingSpace, width))
			}

			fmt.Printf("\n");
		}
	}

	if board.IsLarge {
		fmt.Printf("\n");
	}

	fmt.Printf(center(paddingSpace, width));

	for i, file := range files {
		if !board.IsWhiteSide {
			file = files[(len(files) - 1) - i]
		}

		fmt.Print(center(file, width))
	}

	fmt.Printf("\n");
}

func DrawSquare(isBgWhite bool, piece string, width int) string {
	bgFunc := BgBlack

	if isBgWhite {
		bgFunc = BgWhite
	}

	sprint := fmt.Sprintf("%s", bgFunc(FgBlack(center(piece, width))))

	return sprint
}

var (
	BgBlack	= Color("\033[48;5;248m%s\033[0m")
	BgWhite	= Color("\033[48;5;255m%s\033[0m")
	FgBlack	= Color("\033[38;5;0m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
