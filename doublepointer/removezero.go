package doublepointer

func removeZeroes(nums []int) {
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
}
