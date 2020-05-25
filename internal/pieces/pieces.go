package pieces

import ()

type Piece struct {
	Icon		rune
	IsWhite		bool
	Name		string
	Notation	string
	Position	int
}

var Pieces = map[string]Piece{
	"BK": {
		Icon:		'♚',
		IsWhite:	false,
		Name:		"k",
		Notation:	"K",
	},
	"BQ": {
		Icon:		'♛',
		IsWhite:	false,
		Name:		"q",
		Notation:	"Q",
	},
	"BR": {
		Icon:		'♜',
		IsWhite:	false,
		Name:		"r",
		Notation:	"R",
	},
	"BB": {
		Icon:		'♝',
		IsWhite:	false,
		Name:		"b",
		Notation:	"B",
	},
	"BN": {
		Icon:		'♞',
		IsWhite:	false,
		Name:		"n",
		Notation:	"N",
	},
	"BP": {
		Icon:		'♟',
		IsWhite:	false,
		Name:		"p",
		Notation:	"",
	},
	"WK": {
		Icon:		'♔',
		IsWhite:	true,
		Name:		"K",
		Notation:	"K",
	},
	"WQ": {
		Icon:		'♕',
		IsWhite:	true,
		Name:		"Q",
		Notation:	"Q",
	},
	"WR": {
		Icon:		'♖',
		IsWhite:	true,
		Name:		"R",
		Notation:	"R",
	},
	"WB": {
		Icon:		'♗',
		IsWhite:	true,
		Name:		"B",
		Notation:	"B",
	},
	"WN": {
		Icon:		'♘',
		IsWhite:	true,
		Name:		"N",
		Notation:	"N",
	},
	"WP": {
		Icon:		'♙',
		IsWhite:	true,
		Name:		"P",
		Notation:	"",
	},
}

func GetPiece(fullName string, position int) Piece {
	piece := Pieces[fullName]
	piece.Position = position
	return piece
}

// func validMoves(piece string, position int) func(...interface{}) []int {
// 	switch piece {
// 		case "P":
// 			return func()() {
// 			}
// 	}
// }
