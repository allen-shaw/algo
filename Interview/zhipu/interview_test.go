package zhipu

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(3)
	c.Set(1, 1)
	c.Set(2, 2)
	c.Set(3, 3)
	c.print()

	c.Set(4, 4)
	c.print()

	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	c.Set(5, 5)
	c.print()
}
