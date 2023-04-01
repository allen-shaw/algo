package sort

import (
	"fmt"
	"testing"
)

// 将数组一分为二，左侧排好序，右侧排好序，归并
func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l == r {
		return
	}

	mid := l + (r-l)>>2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)

	merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
	temp := make([]int, r-l+1)

	i, j, k := l, mid+1, 0
	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	for ; i <= mid; i++ {
		temp[k] = arr[i]
		k++
	}

	for ; j <= r; j++ {
		temp[k] = arr[j]
		k++
	}

	for i := 0; i < len(temp); i++ {
		arr[l+i] = temp[i]
	}
}

func TestMerge(t *testing.T) {
	arr := []int{1, 9, 11, 25, 78, 3, 4, 19, 64, 77}
	l, r := 0, len(arr)-1
	mid := l + (r-l)>>1

	merge(arr, l, mid, r)
	fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 1, 9, 7, 6, 4, 3}
	MergeSort(arr)
	fmt.Println(arr)
}

// 时间复杂度
// T(N) = 2 * T(N/2) + O(N)
// a=2,b=2,d=1 logba=d=1
// T(N)=O(N^d*logN)=O(N*logN)
