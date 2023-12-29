package other

import (
	"container/heap"
)

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

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) <= 2 || len(heightMap[0]) <= 2 {
		return 0
	}

	m, n := len(heightMap), len(heightMap[0])
	hp := &hp{}
	visited := newMatrix[bool](m, n)
	res := 0

	// 将周边一圈入队
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[0]); j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(hp, &cell{heightMap[i][j], i, j})
				visited[i][j] = true
			}
		}
	}

	dir := []int{-1, 0, 1, 0, -1}
	for hp.Len() > 0 {
		c := heap.Pop(hp).(*cell)
		for i := 0; i < 4; i++ {
			x, y := c.x+dir[i], c.y+dir[i+1]
			if x >= 0 && x <= m-1 && y >= 0 && y <= n-1 && !visited[x][y] {
				if heightMap[x][y] < c.h {
					res += c.h - heightMap[x][y]
					heap.Push(hp, &cell{c.h, x, y})
				} else {
					heap.Push(hp, &cell{heightMap[x][y], x, y})
				}
				visited[x][y] = true
			}
		}
	}

	return res
}

type cell struct {
	h    int
	x, y int
}
type hp []*cell

// Len implements heap.Interface.
func (h *hp) Len() int {
	return len(*h)
}

// Less implements heap.Interface.
func (h *hp) Less(i int, j int) bool {
	return (*h)[i].h < (*h)[j].h
}

// Pop implements heap.Interface.
func (h *hp) Pop() any {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

// Push implements heap.Interface.
func (h *hp) Push(x any) {
	*h = append(*h, x.(*cell))
}

// Swap implements heap.Interface.
func (h *hp) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func newMatrix[T any](m, n int) [][]T {
	matrix := make([][]T, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]T, n)
	}
	return matrix
}
