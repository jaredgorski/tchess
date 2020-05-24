package pieces

import ()

type Piece struct {
	Name		string
	Notation	string
	Icon		rune
	IsWhite		bool
}

var Pieces = map[string]Piece{
	"BK": {
		Name:		"k",
		Notation:	"K",
		Icon:		'♚',
		IsWhite:	false,
	},
	"BQ": {
		Name:		"q",
		Notation:	"Q",
		Icon:		'♛',
		IsWhite:	false,
	},
	"BR": {
		Name:		"r",
		Notation:	"R",
		Icon:		'♜',
		IsWhite:	false,
	},
	"BB": {
		Name:		"b",
		Notation:	"B",
		Icon:		'♝',
		IsWhite:	false,
	},
	"BN": {
		Name:		"n",
		Notation:	"N",
		Icon:		'♞',
		IsWhite:	false,
	},
	"BP": {
		Name:		"p",
		Notation:	"",
		Icon:		'♟',
		IsWhite:	false,
	},
	"WK": {
		Name:		"K",
		Notation:	"K",
		Icon:		'♔',
		IsWhite:	true,
	},
	"WQ": {
		Name:		"Q",
		Notation:	"Q",
		Icon:		'♕',
		IsWhite:	true,
	},
	"WR": {
		Name:		"R",
		Notation:	"R",
		Icon:		'♖',
		IsWhite:	true,
	},
	"WB": {
		Name:		"B",
		Notation:	"B",
		Icon:		'♗',
		IsWhite:	true,
	},
	"WN": {
		Name:		"N",
		Notation:	"N",
		Icon:		'♘',
		IsWhite:	true,
	},
	"WP": {
		Name:		"P",
		Notation:	"",
		Icon:		'♙',
		IsWhite:	true,
	},
	"_": {
		Name:		" ",
		Icon:		' ',
		IsWhite:	false,
	},
}

// func validMoves(piece string, position int) func(...interface{}) []int {
// 	switch piece {
// 		case "P":
// 			return func()() {
// 			}
// 	}
// }
