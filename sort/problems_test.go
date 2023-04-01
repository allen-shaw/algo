package sort

import (
	"fmt"
	"testing"
)

// func smallSum(arr []int) int {

// }

// 荷兰国旗问题
// 75. 颜色分类
// https://leetcode.cn/problems/sort-colors/
// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
func netherlandsFlag(arr []int) {
	if len(arr) < 2 {
		return
	}

	partitionByKey(arr, 0, len(arr)-1, 1)
}

func partitionByKey(arr []int, l, r int, k int) {
	less, greater := l-1, r+1
	for i := l; i < greater; {
		if arr[i] < k {
			// 当前值小于k，与小于区域的下一位交换，小于区域往右边扩,i++
			arr[i], arr[less+1] = arr[less+1], arr[i]
			less++
			i++
		} else if arr[i] == k {
			// 因为当前在等于区域，所以直接到下一位
			i++
		} else {
			// 当前值大于k,应该与大于区域的前一位交换，大于区域往左扩，i不变
			// 因为交换后的值arr[i]是没看过的，i不能++，等待下一轮
			arr[i], arr[greater-1] = arr[greater-1], arr[i]
			greater--
		}
	}
}

func TestNetherlandsFlag(t *testing.T) {
	arr1 := []int{2, 0, 2, 1, 1, 0}
	netherlandsFlag(arr1)
	fmt.Println(arr1)

	arr2 := []int{2, 0, 1}
	netherlandsFlag(arr2)
	fmt.Println(arr2)
}
