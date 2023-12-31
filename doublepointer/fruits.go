package doublepointer

func totalFruit(fruits []int) int {
	total := 0
	window := make(map[int]int)

	left, right := 0, 0
	for right < len(fruits) {
		fruit := fruits[right]
		window[fruit]++
		right++
		if len(window) <= 2 {
			total = max(right-left, total)
		}
		for left < right && len(window) > 2 {
			fruit := fruits[left]
			window[fruit]--
			if window[fruit] == 0 {
				delete(window, fruit)
			}
			left++
		}
	}

	return total
}
