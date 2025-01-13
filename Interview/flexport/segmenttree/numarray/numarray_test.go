package numarray

import (
	"fmt"
	"testing"
)

type NumArray struct {
	st SegementTree
}

func Constructor(nums []int) NumArray {
	st := NewSegmentTree(nums)
	arr := NumArray{st: st}
	return arr
}

func (a *NumArray) Update(index int, val int) {
	delta := val - a.st.arr[index]
	a.st.Update(index, index, delta)
}

func (a *NumArray) SumRange(left int, right int) int {
	return a.st.Query(left, right)
}

func Test_NumArray(t *testing.T) {
	na := Constructor([]int{1, 3, 5})
	fmt.Println(na.SumRange(0, 2))
	na.Update(1, 2)
	fmt.Println(na.SumRange(0, 2))
}

type SegementTree struct {
	arr   []int
	nodes []int
	lazy  []int
}

func NewSegmentTree(arr []int) SegementTree {
	n := len(arr)

	st := SegementTree{
		arr:   arr,
		nodes: make([]int, n*4),
		lazy:  make([]int, 4*n),
	}
	st.buildTree(0, 0, n-1)

	fmt.Println("st.arr:", st.arr)
	fmt.Println("st.nodes:", st.nodes)
	return st
}

func (s SegementTree) buildTree(nodeId int, start, end int) {
	if start == end {
		s.nodes[nodeId] = s.arr[start]
		return
	}

	mid := (start + end) / 2
	lcid := nodeId*2 + 1
	rcid := nodeId*2 + 2
	s.buildTree(lcid, start, mid)
	s.buildTree(rcid, mid+1, end)

	s.nodes[nodeId] = s.nodes[lcid] + s.nodes[rcid]
}

func (s *SegementTree) Query(left, right int) int {
	return s.query(0, 0, len(s.arr)-1, left, right)
}

func (s *SegementTree) query(nodeId, start, end, left, right int) int {
	if left <= start && end <= right {
		return s.nodes[nodeId]
	}

	s.pushDown(nodeId, start, end)

	ans := 0
	mid := (start + end) / 2
	if left <= mid {
		ans += s.query(nodeId*2+1, start, mid, left, right)
	}
	if right > mid {
		ans += s.query(nodeId*2+2, mid+1, end, left, right)
	}

	return ans
}

func (s *SegementTree) pushDown(nodeId int, start int, end int) {
	if s.lazy[nodeId] == 0 {
		return
	}

	lazy := s.lazy[nodeId]
	lcid, rcid := nodeId*2+1, nodeId*2+2
	mid := (start + end) / 2
	lnum := mid - start + 1
	rnum := end - mid

	s.nodes[lcid] += lazy * lnum
	s.lazy[lcid] += lazy
	s.nodes[rcid] += lazy * rnum
	s.lazy[rcid] += lazy

	s.lazy[nodeId] = 0
}

func (s *SegementTree) Update(left, right, val int) {
	for i := left; i <= right; i++ {
		s.arr[i] += val
	}
	s.update(0, 0, len(s.arr)-1, left, right, val)
}

func (s *SegementTree) update(nodeId, start, end, left, right, val int) {
	if left <= start && end <= right {
		s.nodes[nodeId] += val * (end - start + 1)
		s.lazy[nodeId] += val
		return
	}

	s.pushDown(nodeId, start, end)
	mid := (end + start) / 2
	l, r := nodeId*2+1, nodeId*2+2
	if left <= mid {
		s.update(l, start, mid, left, right, val)
	}
	if right > mid {
		s.update(r, mid+1, end, left, right, val)
	}

	s.nodes[nodeId] = s.nodes[l] + s.nodes[r]
}
