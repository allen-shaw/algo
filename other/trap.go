package other

func trap(height []int) int {
	res := 0
	// preCalcMaxLeft
	maxLefts := preCalcMaxLeft(height)

	// preCalcMaxRight
	maxRights := preCalcMaxRight(height)

	for i := 0; i < len(height); i++ {
		// 计算height[i]左右的最大值
		// maxLeft := calcMaxLeft(height, i)
		// maxRight := calcMaxRight(height, i)
		maxLeft := maxLefts[i]
		maxRight := maxRights[i]

		c := max(min(maxLeft, maxRight)-height[i], 0)
		res += c
	}

	return res
}

// func calcMaxLeft(height []int, idx int) int {
// 	m := 0
// 	for i := idx; i >= 0; i-- {
// 		if height[i] > m {
// 			m = height[i]
// 		}
// 	}
// 	return m
// }

// func calcMaxRight(height []int, idx int) int {
// 	m := 0
// 	for i := idx; i < len(height); i++ {
// 		if height[i] > m {
// 			m = height[i]
// 		}
// 	}
// 	return m
// }

func preCalcMaxLeft(height []int) []int {
	res := make([]int, len(height))
	maxLeft := 0

	for i := 0; i < len(height); i++ {
		if height[i] > maxLeft {
			maxLeft = height[i]
		}
		res[i] = maxLeft
	}

	return res
}

func preCalcMaxRight(height []int) []int {
	res := make([]int, len(height))
	maxRight := 0

	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > maxRight {
			maxRight = height[i]
		}
		res[i] = maxRight
	}

	return res
}

func trap2(height []int) int {
	leftArea := calcLeftArea(height)
	rightArea := calcRightArea(height)
	pillarArea := calcPillarArea(height)
	maxArea := calcMaxArea(height)

	return leftArea + rightArea - maxArea - pillarArea
}

func calcLeftArea(height []int) int {
	res := 0
	maxHegiht := 0
	for i := 0; i < len(height); i++ {
		if height[i] > maxHegiht {
			maxHegiht = height[i]
		}
		res += maxHegiht
	}
	return res
}

func calcRightArea(height []int) int {
	res := 0
	maxHegiht := 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > maxHegiht {
			maxHegiht = height[i]
		}
		res += maxHegiht
	}
	return res
}

func calcPillarArea(height []int) int {
	res := 0
	for _, n := range height {
		res += n
	}
	return res
}

func calcMaxArea(height []int) int {
	maxHegiht := 0
	for _, n := range height {
		maxHegiht = max(maxHegiht, n)
	}
	return len(height) * maxHegiht
}

func trap3(height []int) int {
	res := 0
	lmax, rmax := 0, 0

	l, r := 0, len(height)-1
	for l < r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])

		// res += min(lmax, rmax) - height[i]
		if lmax < rmax {
			res += lmax - height[l]
			l++
		} else {
			res += rmax - height[r]
			r--
		}
	}

	return res
}
