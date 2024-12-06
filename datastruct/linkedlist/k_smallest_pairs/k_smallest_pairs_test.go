package ksmallestpairs

import (
	"container/heap"
	"fmt"
	"testing"
)

// nums1 = [1,7,11], nums2 = [2,4,6]
// 1,2 -> 1,4 -> 1,6
// 7,2 -> 7,4 -> 7,6
// 11,2 -> 11,4 -> 11,6
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	hp := &Heap{}
	heap.Init(hp)
	for i, n1 := range nums1 {
		heap.Push(hp, Pair{i, 0, n1, nums2[0]})
	}

	ans := make([][]int, 0)
	for i := 0; i < k; i++ {
		pair := heap.Pop(hp).(Pair)
		ans = append(ans, []int{pair.x, pair.y})
		x, y := pair.i, pair.j+1
		if y < len(nums2) {
			heap.Push(hp, Pair{x, y, nums1[x], nums2[y]})
		}
	}

	return ans
}

type Pair struct {
	i, j int
	x, y int
}

type Heap []Pair

func (hp *Heap) Push(x any) {
	(*hp) = append((*hp), x.(Pair))
}

func (hp *Heap) Pop() any {
	x := (*hp)[hp.Len()-1]
	(*hp) = (*hp)[:hp.Len()-1]
	return x
}

func (hp Heap) Len() int {
	return len(hp)
}

func (hp Heap) Less(i, j int) bool {
	return hp[i].x+hp[i].y < hp[j].x+hp[j].y
}

func (hp Heap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func Test_kSmallestPairs(t *testing.T) {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3

	out := kSmallestPairs(nums1, nums2, k)
	fmt.Println(out)
}
