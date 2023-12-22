package monostack

// 逆序扫描，使用单调栈（单调递增
// 当元素进入的时候，当前要进入的元素往右可以看到元素就是pop出来的元素+栈顶比当前元素大的元素
func canSeePersonsCount(heights []int) []int {
	res := make([]int, len(heights))
	stack := make([]int, 0)

	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 {
			top := heights[stack[len(stack)-1]]
			if top < heights[i] {
				stack = stack[:len(stack)-1]
				res[i]++
			} else {
				// 加1，代表栈顶元素
				res[i]++
				break
			}
		}
		stack = append(stack, i)
	}

	return res
}
