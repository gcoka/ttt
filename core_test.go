package ttt

import "testing"

func TestNew(t *testing.T) {
	g := New(`
	ooo
	xxx
	___
	`)
	if g.board[0][0] != Circle {
		t.Errorf("g.board[0][0] = %v, want = %v", g.board[0][0], Circle)
	}
	if g.board[1][1] != Cross {
		t.Errorf("g.board[1][1] = %v, want = %v", g.board[1][1], Cross)
	}
	if g.board[2][2] != Empty {
		t.Errorf("g.board[2][2] = %v, want = %v", g.board[2][2], Empty)
	}
	if g.board[2][1] != Cross {
		t.Errorf("g.board[2][1] = %v, want = %v", g.board[2][1], Cross)
	}
}
