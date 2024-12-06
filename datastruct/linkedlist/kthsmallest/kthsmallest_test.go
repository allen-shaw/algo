package kthsmallest

import (
	"container/heap"
	"fmt"
	"testing"
)

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	hp := &Heap{}
	heap.Init(hp)

	ans := 0
	for i := range matrix {
		heap.Push(hp, ELem{i: i, j: 0, val: matrix[i][0]})
	}
	for i := 0; i < k; i++ {
		elem := heap.Pop(hp).(ELem)
		ans = elem.val
		x, y := elem.i, elem.j+1
		if x < n && y < n {
			heap.Push(hp, ELem{i: x, j: y, val: matrix[x][y]})
		}
	}

	return ans
}

type ELem struct {
	i, j int
	val  int
}
type Heap []ELem

func (hp *Heap) Push(x any) {
	(*hp) = append((*hp), x.(ELem))
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
	return hp[i].val < hp[j].val
}

func (hp Heap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func Test_kthSmallest(t *testing.T) {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	k := 8
	fmt.Println(kthSmallest(matrix, k))
}
