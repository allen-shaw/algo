package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getNax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	return doGetMax(arr, 0, len(arr)-1)
}

func doGetMax(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	mid := l + ((r - l) >> 1)
	leftMax := doGetMax(arr, l, mid)
	rightMax := doGetMax(arr, mid+1, r)
	return max(leftMax, rightMax)
}

// T(N) = 2 * T(N/2) + O(1)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestGetMax(t *testing.T) {
	arr := []int{1, 9, 4, 2, 6, 2, 5, 1, 4, 21, 43, 1, 3}
	m := getNax(arr)
	assert.Equal(t, 43, m)
	fmt.Println(m)
}
