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
