package queue

func maxSlidingWindow(nums []int, k int) []int {
	left, right := 0, 0
	mq := NewMonotonicQueue(3)
	res := make([]int, 0)

	for right < len(nums) {
		for right-left < k {
			n := nums[right]
			mq.Push(n)
			right++
		}

		// right - left == k
		for right-left >= k {
			left++
			res = append(res, mq.Max())
			mq.Pop()
		}
	}

	return res
}
