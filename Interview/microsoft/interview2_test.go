package microsoft

import (
	"fmt"
	"testing"
)

// 4，5，1，2，3
func search(arr []int) int {
	n := len(arr)
	if arr[0] < arr[n-1] {
		return arr[0]
	}
	left, right := 0, n-1
	for left < right {
		mid := (right + left) / 2
		if arr[mid] > arr[0] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left]
}

func TestSearch(t *testing.T) {
	arr := []int{4, 5, 6, 2, 3}
	out := search(arr)
	fmt.Println(out)
}

// 二叉树按层遍历
//
//	1
//
// 2 3
// 4 5 6 7
// 1 3 2 4 5 6 7
type Node struct {
	Val         int
	Left, Right *Node
}

func travel(root *Node) []int {
	if root == nil {
		return nil
	}

	ans := make([]int, 0)
	q := []*Node{root}
	odd := 1

	for len(q) > 0 {
		size := len(q)
		temp := make([]int, 0)
		for range size {
			n := q[0]
			q = q[:len(q)-1] // pop
			temp = append(temp, n.Val)

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		if odd > 0 {
			ans = append(ans, temp...)
		} else {
			ans = append(ans, revert(temp)...)
		}
		odd = -odd
	}

	return ans
}

func revert(src []int) []int {
	left, right := 0, len(src)-1
	for left < right {
		src[left], src[right] = src[right], src[left]
		left++
		right--
	}
	return src
}
