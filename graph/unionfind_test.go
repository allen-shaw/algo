package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByte(t *testing.T) {
	g := 'g'
	v := int(g - 'a')
	fmt.Println(v)
}

func TestEquationsPossible(t *testing.T) {
	var flagtests = []struct {
		in  []string
		out bool
	}{
		{[]string{"a==b", "b!=a"}, false},
		{[]string{"b==a", "a==b"}, true},
		{[]string{"a==b", "b==c", "a==c"}, true},
		{[]string{"a==b", "b!=c", "c==a"}, false},
		{[]string{"c==c", "b==d", "x!=z"}, true},
		{[]string{"a!=b", "b!=c", "c!=a"}, true},
		{[]string{"a==d", "f!=b", "f==a", "c!=c"}, false},
		{[]string{"f==d", "b!=e", "d!=c", "b==c", "b!=a", "b!=f"}, true},
	}

	for _, tt := range flagtests {
		t.Run("equationsPossible", func(t *testing.T) {
			out := equationsPossible(tt.in)
			assert.Equal(t, tt.out, out)
		})
	}
}
