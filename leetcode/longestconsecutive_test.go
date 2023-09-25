package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	arr := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(arr))
}

func TestLongestConsecutive2(t *testing.T) {
	arr := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(arr))
}
