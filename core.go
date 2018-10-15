package ttt

import (
	"fmt"
	"strings"
)

// Symbol マルバツの印
type Symbol int

// マルバツの印の値
const (
	Circle Symbol = 1
	Cross  Symbol = -1
	Empty  Symbol = 0
)

// TicTacToe マルバツゲームの情報
type TicTacToe struct {
	board [][]Symbol
}

func New(state string) *TicTacToe {
	s := state
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	g := &TicTacToe{}
	g.board = make([][]Symbol, 3, 3)
	for i := 0; i < 3; i++ {
		g.board[i] = make([]Symbol, 3, 3)
	}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			c := s[x+y*3]
			switch c {
			case 'o':
				g.board[x][y] = Circle
			case 'x':
				g.board[x][y] = Cross
			case '_':
				g.board[x][y] = Empty
			default:
				panic(fmt.Errorf("ttt.New()にo,x,_以外が渡された, got=%v", c))
			}
		}
	}
	return g
}
