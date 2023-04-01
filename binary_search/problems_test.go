package binarysearch

import (
	"fmt"
	"testing"
)

// 162. 寻找峰值
// https://leetcode.cn/problems/find-peak-element
// 峰值元素是指其值严格大于左右相邻值的元素。
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
// 你可以假设 nums[-1] = nums[n] = -∞ 。
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	l, r := 0, len(nums)-1
	for l < r {
		// 先判断0和n-1是不是峰值
		if isPeakElement(nums, l) {
			return l
		}
		if isPeakElement(nums, r) {
			return r
		}

		mid := l + ((r - l) >> 1)
		if isPeakElement(nums, mid) {
			return mid
		} else if nums[mid-1] > nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func isPeakElement(nums []int, i int) bool {
	if i == 0 {
		return nums[i+1] < nums[i]
	}
	n := len(nums)
	if i == n-1 {
		return nums[i-1] < nums[i]
	}

	return nums[i-1] < nums[i] && nums[i] > nums[i+1]
}

func TestFindPeakElement(t *testing.T) {
	// nums1 := []int{1, 2, 3, 1}
	// k1 := findPeakElement(nums1)
	// fmt.Println(k1)

	// nums2 := []int{1, 2, 1, 3, 5, 6, 4}
	// k2 := findPeakElement(nums2)
	// fmt.Println(k2)

	// nums3 := []int{3, 4, 3, 2, 1}
	// k3 := findPeakElement(nums3)
	// fmt.Println(k3)

	nums4 := []int{1, 2, 3, 1}
	k4 := findPeakElement(nums4)
	fmt.Println(k4)
}

// 852. 山脉数组的峰顶索引
// https://leetcode.cn/problems/peak-index-in-a-mountain-array/
// 符合下列属性的数组 arr 称为 山脉数组 ：
// arr.length >= 3
// 存在 i（0 < i < arr.length - 1）使得：
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
// 给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。

func PeakIndexInMountainArray(arr []int) int {
	if len(arr) < 3 {
		return -1
	}

	l, r := 1, len(arr)-2
	// 判断arr[1] 和 arr[len(arr)-2]是不是peek

	return peakIndexInMountainArray(arr, l, r)
}

func peakIndexInMountainArray(arr []int, l, r int) int {
	if l == r {
		return l
	}
	if isPeek(arr, l) {
		return l
	}
	if isPeek(arr, r) {
		return r
	}

	mid := l + (r-l)>>1

	if isPeek(arr, mid) {
		return mid
	}

	if arr[mid-1] < arr[mid] {
		return peakIndexInMountainArray(arr, mid+1, r)
	} else {
		return peakIndexInMountainArray(arr, l, mid)
	}
}

func isPeek(arr []int, i int) bool {
	if i == 0 || i == len(arr)-1 {
		return false
	}

	return arr[i-1] < arr[i] && arr[i] > arr[i+1]
}

func TestPeakIndexInMountainArray(t *testing.T) {
	arr1 := []int{0, 1, 0}
	r1 := PeakIndexInMountainArray(arr1)
	fmt.Println(r1)

	arr2 := []int{0, 2, 1, 0}
	r2 := PeakIndexInMountainArray(arr2)
	fmt.Println(r2)

	arr3 := []int{0, 10, 5, 2}
	r3 := PeakIndexInMountainArray(arr3)
	fmt.Println(r3)

	arr4 := []int{3, 4, 5, 1}
	r4 := PeakIndexInMountainArray(arr4)
	fmt.Println(r4)

	arr5 := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	r5 := PeakIndexInMountainArray(arr5)
	fmt.Println(r5)
}
