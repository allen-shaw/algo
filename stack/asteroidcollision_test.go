package stack

import (
	"fmt"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	asts := []int{-5, 10, -5, 8, -9}
	ans := asteroidCollision(asts)
	t.Log(ans)
}

func TestLoopBreak(t *testing.T) {
	for i := 0; i < 5; i++ {
	loop:
		for j := 0; j < 10; j++ {
			fmt.Printf("%v:%v ", i, j)
			if i == j {
				break loop
			}
		}
		fmt.Println()
	}
}
