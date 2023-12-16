package nsums

import "testing"

func TestTwoSumBsts(t *testing.T) {
	root1 := buildTree([]int{2, 1, 4})
	root2 := buildTree([]int{1, 0, 3})
	target := 5

	ans := twoSumBsts(root1, root2, target)
	t.Log(ans)
}

// root1 = [0,-10,10], root2 = [5,1,7,0,2], target = 18

func TestTwoSumBsts2(t *testing.T) {
	root1 := buildTree([]int{0, -10, 10})
	root2 := buildTree([]int{5, 1, 7, 0, 2})
	target := 18

	ans := twoSumBsts(root1, root2, target)
	t.Log(ans)
}
