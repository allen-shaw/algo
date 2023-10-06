package queue

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	ans := maxSlidingWindow(nums, k)
	fmt.Println(ans)
}

func TestMaxSlidingWindow2(t *testing.T) {
	nums := []int{1}
	k := 1
	ans := maxSlidingWindow(nums, k)
	fmt.Println(ans)
}

func TestMaxSlidingWindow3(t *testing.T) {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3
	ans := maxSlidingWindow(nums, k)
	fmt.Println(ans)
}

func TestMaxSlidingWindow4(t *testing.T) {
	nums := []int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}
	k := 7
	ans := maxSlidingWindow(nums, k)
	fmt.Println(ans)
}
