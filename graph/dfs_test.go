package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaRound(t *testing.T) {
	var testset = []struct {
		in  [][]byte
		out [][]byte
	}{
		// {
		// 	[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
		// 	[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		// },
		{
			[][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}},
			[][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}},
		},
	}

	for _, tt := range testset {
		t.Run("areaRound", func(t *testing.T) {
			areaRound(tt.in)
			assert.Equal(t, tt.out, tt.in)
		})
	}
}

func TestExist(t *testing.T) {
	var testsets = []struct {
		board [][]byte
		word  string
		ans   bool
	}{
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCCED",
			true,
		},
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"SEE",
			true,
		},
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCB",
			false,
		},
	}

	for _, tt := range testsets {
		t.Run("exist", func(t *testing.T) {
			ans := exist(tt.board, tt.word)
			assert.Equal(t, tt.ans, ans)
		})
	}
}
