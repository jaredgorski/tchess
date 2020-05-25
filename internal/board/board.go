package board

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/jaredgorski/tchess/internal/pieces"
)

// move to util
var paddingSpace = string('\u00A0')

type Square struct {
	IsHighlighted	bool
	Piece			pieces.Piece
}

// Board represents the chess board.
//	Squares is an array that indexes each square of the chess board from
//	0 (A1) to 63 (H8).
type Board struct {
	IsWhiteSide	bool
	IsLarge		bool
	UseIcons	bool
	LastSquare	int
	Squares		[64]Square
}

// move to util
func center(s string, n int) string {
	div := n / 2

	return strings.Repeat(paddingSpace, div) + s + strings.Repeat(paddingSpace, div)
}

var filesIndices = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
}

var files = [8]string{"A", "B", "C", "D", "E", "F", "G", "H"};
var ranks = [8]string{"1", "2", "3", "4", "5", "6", "7", "8"};

func (board *Board) setPieces(fullName string, positions []int) {
	for _, position := range positions {
		board.Squares[position] = Square{
			IsHighlighted: false,
			Piece: pieces.GetPiece(fullName, position),
		}
	}
}

func (board *Board) ResetBoard() {
	indices := make([]int, 64)

	for i := 0; i < 64; i++ {
		indices[i] = i
	}

	nearColor	:= "W"
	farColor	:= "B"

	if !board.IsWhiteSide {
		nearColor	= "B"
		farColor	= "W"
	}

	// pawns : 8 - 16 , 48 - 56
	board.setPieces(farColor + "P", indices[8:16])
	board.setPieces(nearColor + "P", indices[48:56])

	// rooks : 0 , 7 , 56 , 63
	board.setPieces(farColor + "R", indices[0:1])
	board.setPieces(farColor + "R", indices[7:8])
	board.setPieces(nearColor + "R", indices[56:57])
	board.setPieces(nearColor + "R", indices[63:64])

	// knights : 1 , 6 , 57 , 62
	board.setPieces(farColor + "N", indices[1:2])
	board.setPieces(farColor + "N", indices[6:7])
	board.setPieces(nearColor + "N", indices[57:58])
	board.setPieces(nearColor + "N", indices[62:63])

	// bishops : 2 , 5 , 58 , 61
	board.setPieces(farColor + "B", indices[2:3])
	board.setPieces(farColor + "B", indices[5:6])
	board.setPieces(nearColor + "B", indices[58:59])
	board.setPieces(nearColor + "B", indices[61:62])

	// queens : 3 , 59
	board.setPieces(farColor + "Q", indices[3:4])
	board.setPieces(nearColor + "Q", indices[59:60])

	// kings : 4 , 60
	board.setPieces(farColor + "K", indices[4:5])
	board.setPieces(nearColor + "K", indices[60:61])
}

func (board Board) MovePiece(move string) {
	// - get last two characters of move and convert to new position
	// - check if "x" to see if piece is taken on new position
	// - if more than two characters left in string, use previous two characters as old position
	//	- else if only two characters left in string, error
	//	- else if one character left in string, use character as piece
	//	- else, find pawn to advance

	moveSlice := strings.Split(move, "")

	var (
		newPosCoord	string
		newPos		int
		oldPosCoord	string
		oldPos		int
		pieceName	string
	)

	for len(moveSlice) > 0 {
		// pop last character from moveSlice
		lastIndex := len(moveSlice) - 1
		char := moveSlice[lastIndex]
		moveSlice[lastIndex] = ""
		moveSlice = moveSlice[:lastIndex]

		if len(newPosCoord) < 1 {
			i, err := strconv.Atoi(char)
			if err != nil {
			}

			newPosCoord = char
			newPos = (8 - i) * 8

			continue
		} else if len(newPosCoord) < 2 {
			newPosCoord = char + newPosCoord
			newPos += filesIndices[strings.ToUpper(char)]

			continue
		}

		if len(oldPosCoord) < 1 {
			i, err := strconv.Atoi(char)
			if err != nil {
			}

			oldPosCoord = char
			oldPos = (8 - i) * 8

			continue
		} else if len(oldPosCoord) < 2 {
			oldPosCoord = char + oldPosCoord
			oldPos += filesIndices[strings.ToUpper(char)]

			continue
		}

		if len(pieceName) < 1 {
			pieceName = char
		}
	}

	// once parsed, execute move

	fmt.Print(pieceName, oldPosCoord, newPosCoord)
	fmt.Println("")
	fmt.Print(pieceName, oldPos, newPos)

	piece := board.Squares[oldPos].Piece
	board.Squares[oldPos].Piece = pieces.Piece{}
	board.Squares[newPos].Piece = piece

	board.DrawBoard()
}

func (board Board) DrawBoard() {
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

			for j := 0; j < len(files); j++ {
				flag = !flag

				fmt.Printf("%v", DrawSquare(flag, paddingSpace, width))
			}

			fmt.Printf("\n");
		}

		fmt.Print(center(rank, width))

		for j := 0; j < len(files); j++ {
			flag = !flag

			index := (i * 8) + j
			square := board.Squares[index]
			icon := square.Piece.Name

			if len(icon) > 0 {
				if board.UseIcons {
					icon = string(square.Piece.Icon)
				}
			} else {
				icon = " "
			}

			fmt.Printf("%v", DrawSquare(flag, icon, width))
		}

		fmt.Printf("\n");

		if board.IsLarge {
			fmt.Print(center(paddingSpace, width))

			for j := 0; j < len(files); j++ {
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

// move to util
var (
	BgBlack	= Color("\033[48;5;248m%s\033[0m")
	BgWhite	= Color("\033[48;5;255m%s\033[0m")
	FgBlack	= Color("\033[38;5;0m%s\033[0m")
)

// move to util
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
