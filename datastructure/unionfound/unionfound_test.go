package unionfound

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFound(t *testing.T) {
	var testsets = []struct {
		in  []int
		out bool
	}{
		{[]int{0, 1}, true},
	}

	uf := New(2)

	for _, tt := range testsets {
		uf.Union(tt.in[0], tt.in[1])
		out := uf.IsConnected(tt.in[0], tt.in[1])
		assert.Equal(t, tt.out, out)
	}
}
