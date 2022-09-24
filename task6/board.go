package main

import (
	"strings"
	"utils"
)

type BoardCell uint8

const (
	X_SET BoardCell = iota
	O_SET
	EMPTY
)

const (
	BoardSize    = 3
	WinningScore = 10
	LosingScore  = -10
)

type Board [BoardSize][BoardSize]BoardCell

func NewBoard() *Board {
	return &Board{
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
	}
}

func (b *Board) PlayX(coord utils.Coord) {
	b[coord.X][coord.Y] = X_SET
}

func (b *Board) PlayO(coord utils.Coord) {
	b[coord.X][coord.Y] = O_SET
}

func (b *Board) CanPlay(coord utils.Coord) bool {
	return b[coord.X][coord.Y] == EMPTY
}

func (b *Board) IsTieGame() bool {
	return b[0][0] != EMPTY && b[0][1] != EMPTY && b[0][2] != EMPTY &&
		b[1][0] != EMPTY && b[1][1] != EMPTY && b[1][2] != EMPTY &&
		b[2][0] != EMPTY && b[2][1] != EMPTY && b[2][2] != EMPTY
}

func (b *Board) IsGameOver() bool {
	return b.IsWinningBoardFor(X_SET) || b.IsWinningBoardFor(O_SET) || b.IsTieGame()
}

func (b *Board) GetScore() int {
	if b.IsWinningBoardFor(X_SET) {
		return LosingScore
	} else if b.IsWinningBoardFor(O_SET) {
		return WinningScore
	} else {
		return 0
	}
}

func (b *Board) GetPossibleMovesForPlayer(cell BoardCell) []*Board {
	res := make([]*Board, 0, 8)
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if b[i][j] == EMPTY {
				childBoard := *b
				childBoard[i][j] = cell
				res = append(res, &childBoard)
			}
		}
	}
	return res
}

func (b *Board) String() string {
	builder := strings.Builder{}
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			switch b[i][j] {
			case O_SET:
				builder.WriteByte('O')
			case X_SET:
				builder.WriteByte('X')
			case EMPTY:
				builder.WriteByte('_')
			}
			if j < BoardSize-1 {
				builder.WriteString("\t")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (b *Board) IsWinningBoardFor(sign BoardCell) bool {
	return (b[0][0] == sign && b[0][1] == sign && b[0][2] == sign) ||
		(b[1][0] == sign && b[1][1] == sign && b[1][2] == sign) ||
		(b[2][0] == sign && b[2][1] == sign && b[2][2] == sign) ||
		(b[0][0] == sign && b[1][0] == sign && b[2][0] == sign) ||
		(b[0][1] == sign && b[1][1] == sign && b[2][1] == sign) ||
		(b[0][2] == sign && b[1][2] == sign && b[2][2] == sign) ||
		(b[0][0] == sign && b[1][1] == sign && b[2][2] == sign) ||
		(b[0][2] == sign && b[1][1] == sign && b[2][0] == sign)
}
