package monostack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums2))

	// 计算nums2中下一个更大的值
	// 使用单调递增栈
	stack := make([]int, 0, len(nums2))

	for i := 0; i < len(nums2); i++ {
		n := nums2[i]
		for len(stack) > 0 && n > stack[len(stack)-1] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			m[top] = n
		}
		stack = append(stack, n)
		m[n] = -1
	}

	res := make([]int, len(nums1))
	for i, n := range nums1 {
		res[i] = m[n]
	}

	return res
}

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	stack := make([]int, 0, len(nums))

	for i := 0; i < 2*len(nums); i++ {
		j := i % len(nums)
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[j] {
			// pop
			res[stack[len(stack)-1]] = nums[j]
			stack = stack[:len(stack)-1]
		}

		// push
		stack = append(stack, j)
	}

	return res
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		cur := temperatures[i]

		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < cur {
			topIdx := stack[len(stack)-1]
			res[topIdx] = i - topIdx
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
