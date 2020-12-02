package board

import (
	"fmt"
	"io"
	"strings"
	"strconv"

	"github.com/jaredgorski/tchess/internal/pieces"
)

// TODO: move to util
var paddingSpace = string('\u00A0')

type Square struct {
	IsHighlighted	bool
	Piece			pieces.Piece
}

// Board represents the chess board.
//	Squares is an array that indexes each square of the chess board from
//	0 (A1) to 63 (H8).
type Board struct {
	IsWhiteSide		bool
	IsLarge			bool
	IconType		string
	LastSquare		int
	Squares			[64]Square
	Writer			io.Writer
}

// TODO: move to util
func center(s string, n int) string {
	div := n / 2

	return strings.Repeat(paddingSpace, div) + s + strings.Repeat(paddingSpace, div)
}

var filesIndicesWhite = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
}

var filesIndicesBlack = map[string]int{
	"A": 7,
	"B": 6,
	"C": 5,
	"D": 4,
	"E": 3,
	"F": 2,
	"G": 1,
	"H": 0,
}

var files = [8]string{"A", "B", "C", "D", "E", "F", "G", "H"};
var ranks = [8]string{"1", "2", "3", "4", "5", "6", "7", "8"};

func (board *Board) setPieces(fullName string, positions []int) {
	for _, position := range positions {
		board.Squares[position] = Square{
			IsHighlighted: false,
		}

		if len(fullName) > 0 {
			board.Squares[position].Piece = pieces.GetPiece(fullName, position)
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

	// queens : 3 , 59 ; 4, 60
	if board.IsWhiteSide {
		board.setPieces(farColor + "Q", indices[3:4])
		board.setPieces(nearColor + "Q", indices[59:60])
	} else {
		board.setPieces(farColor + "Q", indices[4:5])
		board.setPieces(nearColor + "Q", indices[60:61])
	}

	// kings : 4 , 60 ; 3, 59
	if board.IsWhiteSide {
		board.setPieces(farColor + "K", indices[4:5])
		board.setPieces(nearColor + "K", indices[60:61])
	} else {
		board.setPieces(farColor + "K", indices[3:4])
		board.setPieces(nearColor + "K", indices[59:60])
	}

	// empty squares : 17 - 47
	board.setPieces("", indices[17:47])
}

func (board Board) ParseMove(move string) (int, int) {
	// - get last two characters of move and convert to new position
	// - check if "x" to see if piece is taken on new position
	// - if more than two characters left in string, use previous two characters as old position
	//	- else if only two characters left in string, error
	//	- else if one character left in string, use character as piece
	//	- else, find pawn to advance

	moveSlice := strings.Split(move, "")

	var (
		filesIndices	map[string]int
		newPosCoord		string
		newPos			int
		oldPosCoord		string
		oldPos			int
		pieceName		string
	)

	if board.IsWhiteSide {
		filesIndices = filesIndicesWhite
	} else {
		filesIndices = filesIndicesBlack
	}

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

			if board.IsWhiteSide {
				newPos = (8 - i) * 8
			} else {
				newPos = (i - 1) * 8
			}

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

			if board.IsWhiteSide {
				oldPos = (8 - i) * 8
			} else {
				oldPos = (i - 1) * 8
			}

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

	return oldPos, newPos
}

func (board *Board) MovePiece(oldPos int, newPos int) {
	piece := board.Squares[oldPos].Piece
	board.Squares[oldPos].Piece = pieces.Piece{}
	board.Squares[newPos].Piece = piece
}

func (board Board) IsValidMove(oldPos int, newPos int) bool {
	piece := board.Squares[oldPos].Piece
	validPieceMove := piece.ValidPieceMove(oldPos, newPos)
	validBoardMove := newPos < 64
	validNewSquare := board.Squares[newPos].Piece.Name == "" || board.Squares[newPos].Piece.IsWhite != piece.IsWhite

	isValid := validPieceMove && validBoardMove && validNewSquare
	return isValid
}

func (board Board) DrawBoard() string {
	flag := true
	width := 2

	if board.IsLarge {
		width = 5
	}

	out := ""
	out += "\n";

	for i, rank := range ranks {
		flag = !flag

		if board.IsWhiteSide {
			rank = ranks[(len(ranks) - 1) - i]
		}

		if board.IsLarge {
			out += center(paddingSpace, width)

			for j := 0; j < len(files); j++ {
				flag = !flag

				out += fmt.Sprintf("%v", GenerateSquare(flag, false, paddingSpace, width))
			}

			out += "\n"
		}

		out += center(rank, width)

		for j := 0; j < len(files); j++ {
			flag = !flag

			index := (i * 8) + j
			square := board.Squares[index]
			piece := square.Piece
			icon := piece.Name

			if len(icon) > 0 {
				if board.IconType == "outline" {
					icon = string(square.Piece.IconOutline)
				} else if board.IconType == "filled" {
					icon = string(square.Piece.Icon)
				}
			} else {
				icon = " "
			}

			out += fmt.Sprintf("%v", GenerateSquare(flag, piece.IsWhite, icon, width))
		}

		out += "\n"

		if board.IsLarge {
			out += center(paddingSpace, width)

			for j := 0; j < len(files); j++ {
				flag = !flag

				out += fmt.Sprintf("%v", GenerateSquare(flag, false, paddingSpace, width))
			}

			out += "\n"
		}
	}

	if board.IsLarge {
		out += "\n"
	}

	out += center(paddingSpace, width);

	for i, file := range files {
		if !board.IsWhiteSide {
			file = files[(len(files) - 1) - i]
		}

		out += center(file, width)
	}

	out += "\n"

	return out
}

func GenerateSquare(isBgWhite bool, isFgWhite bool, piece string, width int) string {
	bgFunc := BgBlack

	if isBgWhite {
		bgFunc = BgWhite
	}

	fgFunc := FgBlack

	if isFgWhite {
		fgFunc = FgWhite
	}

	sprint := fmt.Sprintf("%s", bgFunc(fgFunc(center(piece, width))))

	return sprint
}

// TODO: move to util
var (
	BgBlack	= Color("\033[48;5;137m%s\033[0m")
	BgWhite	= Color("\033[48;5;180m%s\033[0m")
	FgBlack	= Color("\033[38;5;0m%s\033[0m")
	FgWhite	= Color("\033[38;5;255m%s\033[0m")
)

// TODO: move to util
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
