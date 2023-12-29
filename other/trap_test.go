package other

import "testing"

func TestTrap(t *testing.T) {
	// h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	h := []int{4, 2, 0, 3, 2, 5}
	ans := trap(h)
	t.Log(ans)
}

func TestPreCalcMaxLeft(t *testing.T) {
	h := []int{4, 2, 0, 3, 2, 5}
	ans := preCalcMaxLeft(h)
	t.Log(ans)
}

func TestTrap2(t *testing.T) {
	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// h := []int{4, 2, 0, 3, 2, 5}
	ans := trap2(h)
	t.Log(ans)
}

func TestTrap3(t *testing.T) {
	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// h := []int{4, 2, 0, 3, 2, 5}
	ans := trap3(h)
	t.Log(ans)
}

func TestTrapRainWater(t *testing.T) {
	heightMap := [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}
	ans := trapRainWater(heightMap)
	t.Log(ans)
}

func TestTrapRainWater2(t *testing.T) {
	heightMap := [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}
	ans := trapRainWater(heightMap)
	t.Log(ans)
}
