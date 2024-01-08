package cache

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	nums := []int{7, 4, 2, 8, 0, 10}
	h := newHeap()
	h.Init(nums)

	// fmt.Println("inited", h.data)

	ans := make([]int, 0)
	for i := 0; i < 4; i++ {
		ans = append(ans, h.Pop())
	}
	fmt.Println(ans) // 0,2,4,7
	ans = make([]int, 0)

	h.Push(3)
	h.Push(5)
	h.Push(9)
	h.Push(100)

	for h.Len() > 0 {
		ans = append(ans, h.Pop())
	}
	fmt.Println(ans) // 3, 5, 8, 9, 10, 100
}
