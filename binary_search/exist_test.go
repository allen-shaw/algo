package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BinarySearch(arr []int, v int) int {
	if len(arr) == 0 {
		return -1
	}

	l, r := 0, len(arr)-1
	for l < r {
		mid := l + ((r - l) >> 1)
		if arr[mid] == v {
			return mid
		} else if arr[mid] < v {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if arr[l] == v {
		return l
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	arr1 := []int{1, 3, 5, 5, 9, 12, 43, 44, 44, 59}
	i := BinarySearch(arr1, 9)
	assert.Equal(t, 4, i)

	i = BinarySearch(arr1, 13)
	assert.Equal(t, -1, i)
}
