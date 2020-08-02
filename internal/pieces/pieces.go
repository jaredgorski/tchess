package pieces

import (
	"math"
)

type validPieceMoveFn func(int, int) bool

type Piece struct {
	Icon			rune
	IsWhite			bool
	Name			string
	Notation		string
	Position		int
	ValidPieceMove	validPieceMoveFn
}

var Pieces = map[string]Piece{
	"BK": {
		Icon:		'♚',
		IsWhite:	false,
		Name:		"k",
		Notation:	"K",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["K"](oldPos, newPos)
		},
	},
	"BQ": {
		Icon:		'♛',
		IsWhite:	false,
		Name:		"q",
		Notation:	"Q",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["Q"](oldPos, newPos)
		},
	},
	"BR": {
		Icon:		'♜',
		IsWhite:	false,
		Name:		"r",
		Notation:	"R",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["R"](oldPos, newPos)
		},
	},
	"BB": {
		Icon:		'♝',
		IsWhite:	false,
		Name:		"b",
		Notation:	"B",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["B"](oldPos, newPos)
		},
	},
	"BN": {
		Icon:		'♞',
		IsWhite:	false,
		Name:		"n",
		Notation:	"N",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["N"](oldPos, newPos)
		},
	},
	"BP": {
		Icon:		'♟',
		IsWhite:	false,
		Name:		"p",
		Notation:	"",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["BP"](oldPos, newPos)
		},
	},
	"WK": {
		Icon:		'♔',
		IsWhite:	true,
		Name:		"K",
		Notation:	"K",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["K"](oldPos, newPos)
		},
	},
	"WQ": {
		Icon:		'♕',
		IsWhite:	true,
		Name:		"Q",
		Notation:	"Q",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["Q"](oldPos, newPos)
		},
	},
	"WR": {
		Icon:		'♖',
		IsWhite:	true,
		Name:		"R",
		Notation:	"R",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["R"](oldPos, newPos)
		},
	},
	"WB": {
		Icon:		'♗',
		IsWhite:	true,
		Name:		"B",
		Notation:	"B",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["B"](oldPos, newPos)
		},
	},
	"WN": {
		Icon:		'♘',
		IsWhite:	true,
		Name:		"N",
		Notation:	"N",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["N"](oldPos, newPos)
		},
	},
	"WP": {
		Icon:		'♙',
		IsWhite:	true,
		Name:		"P",
		Notation:	"",
		ValidPieceMove: func(oldPos int, newPos int) bool {
			return validPieceMoves["WP"](oldPos, newPos)
		},
	},
}

var validPieceMoves = map[string]validPieceMoveFn{
	"K": func(oldPos int, newPos int) bool {
		var absoluteDifference int = int(math.Abs(float64(newPos - oldPos)))

		var isValid bool = absoluteDifference == 1
		return isValid;
	},
	"Q": func(oldPos int, newPos int) bool {
		var absoluteDifference int = int(math.Abs(float64(newPos - oldPos)))
		var isFileOrRankWise bool = absoluteDifference % 8 == 0
		var isDiagonalNeg bool = absoluteDifference % 7 == 0
		var isDiagonalPos bool = absoluteDifference % 9 == 0

		var isValid bool = isFileOrRankWise || isDiagonalNeg || isDiagonalPos
		return isValid;
	},
	"R": func(oldPos int, newPos int) bool {
		var absoluteDifference int = int(math.Abs(float64(newPos - oldPos)))
		var isFileOrRankWise bool = absoluteDifference % 8 == 0

		var isValid bool = isFileOrRankWise
		return isValid;
	},
	"B": func(oldPos int, newPos int) bool {
		var absoluteDifference int = int(math.Abs(float64(newPos - oldPos)))
		var isDiagonalNeg bool = absoluteDifference % 7 == 0
		var isDiagonalPos bool = absoluteDifference % 9 == 0

		var isValid bool = isDiagonalNeg || isDiagonalPos
		return isValid;
	},
	"N": func(oldPos int, newPos int) bool {
		var isValid bool = true
		return isValid;
	},
	"BP": func(oldPos int, newPos int) bool {
		var isValid bool = true
		return isValid;
	},
	"WP": func(oldPos int, newPos int) bool {
		var isValid bool = true
		return isValid;
	},
}

func GetPiece(fullName string, position int) Piece {
	piece := Pieces[fullName]
	piece.Position = position
	return piece
}
