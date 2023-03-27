package sort

import (
	"fmt"
	"testing"
)

// k从0到n-1,分别保证0到k有序
// 类似扑克牌，每轮选第k位数，往前遍历，k比k-1小则交换,类似将k插入到合适的位置
func InsertSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// 从0到n-1,分别使其有序
	for i := 0; i < len(arr); i++ {
		sort(arr, 0, i)
	}
}

func sort(arr []int, i, j int) {
	for k := j - 1; k >= i; k-- {
		if arr[k] > arr[k+1] {
			swap(arr, k, k+1)
		} else {
			break
		}
	}
}

func TestSort(t *testing.T) {
	arr := []int{1, 3, 9, 7, 2}
	sort(arr, 0, 4)
	fmt.Println(arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{1, 3, 9, 7, 2}
	InsertSort(arr)
	fmt.Println(arr)
}
