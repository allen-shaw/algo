package backtrack

import "testing"

func TestSolveNQueen(t *testing.T) {
	t.Log(solveNQueens(1))
	t.Log(solveNQueens(4))
}
