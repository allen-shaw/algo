package binarysearch

import "testing"

// 在arr上，找满足>=value的最左位置
func NearestIndex(arr []int, v int) int {
	if len(arr) == 0 {
		return -1
	}

	idx := -1
	l, r := 0, len(arr)-1
	for l < r {
		mid := l + ((r - l) >> 1)
		if arr[mid] < v {
			l = mid + 1
		} else {
			idx = mid
			r = mid - 1
		}
	}

	return idx
}

func TestNearestIndex(t *testing.T) {

}
