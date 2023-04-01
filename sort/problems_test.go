package sort

import (
	"fmt"
	"testing"
)

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

func smallSum(arr []int, l, r int) int {
	if l >= r {
		return 0
	}

	// 对arr[l,r]进行归并排序
	mid := l + (r-l)>>1
	ls := smallSum(arr, l, mid)
	rs := smallSum(arr, mid+1, r)

	return ls + rs + mergeSum(arr, l, r, mid)
}

func mergeSum(arr []int, l, r, mid int) int {
	// 在arr[l,mid], arr[mid+1,r]进行merge
	sum := 0
	temp := make([]int, r-l+1)
	i, j, k := l, mid+1, 0
	for i <= mid && j <= r {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			sum += (r - j + 1) * arr[i]
			k++
			i++
		} else if arr[i] == arr[j] {
			temp[k] = arr[j]
			k++
			j++
		} else {
			temp[k] = arr[j]
			j++
			k++
		}
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= r {
		temp[k] = arr[j]
		j++
		k++
	}

	for n, v := range temp {
		arr[l+n] = v
	}

	return sum
}

func TestSmallSum(t *testing.T) {
	nums := []int{-2, 5, -1}
	lower, upper := 0, len(nums)-1

	sum := smallSum(nums, lower, upper)
	fmt.Println(sum, nums)
}

// 逆序对
func InversionCount(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	return inversionCount(arr, 0, len(arr)-1)
}

func inversionCount(arr []int, l, r int) int {
	if l == r {
		return 0
	}

	mid := l + (r-l)>>1

	lc := inversionCount(arr, l, mid)
	rc := inversionCount(arr, mid+1, r)
	return lc + rc + mergeCount(arr, l, r, mid)
}

// 找左侧大于右侧的数, arr[l:mid] 和 arr[mid+1,r]都是有序的（逆序）
func mergeCount(arr []int, l, r, mid int) int {
	count := 0
	temp := make([]int, r-l+1)
	i, j, k := l, mid+1, 0
	for i <= mid && j <= r {
		if arr[i] < arr[j] {
			temp[k] = arr[j]
			j++
			k++
		} else if arr[i] == arr[j] {
			temp[k] = arr[j]
			j++
			k++
		} else { // arr[i] > arr[j]
			temp[k] = arr[i]
			i++
			k++
			count += (arr[r] - arr[j] + 1)
		}
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= r {
		temp[k] = arr[j]
		j++
		k++
	}

	for i, v := range temp {
		arr[l+i] = v
	}

	return count
}

func TestInversionCount(t *testing.T) {
	nums := []int{5, 2, 6, 1}
	count := InversionCount(nums)
	fmt.Println(count, nums)
}

// 315. 计算右侧小于当前元素的个数
// https://leetcode.cn/problems/count-of-smaller-numbers-after-self/
// 给你一个整数数组 nums ，按要求返回一个新数组 counts 。
// 数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。
func CountSmaller(nums []int) []int {
	if len(nums) < 2 {
		return []int{0}
	}

	arr := buildNumbers(nums)
	counts := make([]int, len(nums))
	countSmaller(arr, 0, len(arr)-1, counts)

	fmt.Println(arr, counts)

	return counts
}

type Number struct {
	Val   int
	Index int
}

func buildNumbers(arr []int) []Number {
	nums := make([]Number, len(arr))
	for i, v := range arr {
		nums[i] = Number{Val: v, Index: i}
	}
	return nums
}

func countSmaller(arr []Number, l, r int, counts []int) {
	if l == r {
		return
	}

	mid := l + (r-l)>>1
	countSmaller(arr, l, mid, counts)
	countSmaller(arr, mid+1, r, counts)
	mergeCounts(arr, l, r, mid, counts)
}

func mergeCounts(arr []Number, l, r, mid int, counts []int) {
	temp := make([]Number, r-l+1)
	i, j, k := l, mid+1, 0

	for i <= mid && j <= r {
		if arr[i].Val < arr[j].Val {
			temp[k] = arr[j]
			j++
			k++
		} else if arr[i].Val == arr[j].Val {
			temp[k] = arr[j]
			j++
			k++
		} else { // arr[i] > arr[j]
			counts[arr[i].Index] += r - j + 1
			temp[k] = arr[i]
			i++
			k++
		}
	}

	for i <= mid {
		temp[k] = arr[i]
		k++
		i++
	}

	for j <= r {
		temp[k] = arr[j]
		k++
		j++
	}

	for i, v := range temp {
		arr[l+i] = v
	}
}

func TestCountSmaller(t *testing.T) {
	// nums := []int{5, 2, 6, 1}
	nums := []int{-1, -1}
	counts := CountSmaller(nums)
	fmt.Println(counts)
}
