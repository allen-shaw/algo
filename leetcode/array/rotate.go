package array

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)

	fmt.Println("1:", nums)

	reverse(nums, 0, k-1)
	fmt.Println("2:", nums)

	reverse(nums, k, len(nums)-1)
	fmt.Println("3:", nums)

}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
