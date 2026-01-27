package kody

import (
	"fmt"
	"testing"
)

// array: [4, 5, 6, 1, 2, 3]
// value: 3
// Output: 5

func bisect(arr []int, target int) int {
	n := len(arr)
	left, right := 0, n-1

	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == target {
			return mid
		}
		if target <= arr[n-1] {
			if arr[mid] > arr[n-1] || arr[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if arr[mid] > arr[n-1] {
				if arr[mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

//array: [33, 45, 55, 67, 98, 8, 11, 19, 23, 27]
// value: 10
// Output: -1

func Test_bisect(t *testing.T) {
	// arr := []int{4, 5, 6, 1, 2, 3}
	arr := []int{33, 45, 55, 67, 98, 8, 11, 19, 23, 27}
	// arr := []int{1, 2, 3, 4, 5}
	ans := bisect(arr, -1)
	fmt.Println(ans)
}
