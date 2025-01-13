package flexport

import (
	"fmt"
	"testing"
)

type Node struct {
	lazy        int
	val         int
	left, right *Node
}

func (n *Node) Query(start, end, left, right int) int {
	// fmt.Println("Query:", start, end, left, right)
	if left <= start && right >= end {
		if n.val > 0 {
			fmt.Println("Query:", start, end, left, right, n.val)
		}
		return n.val
	}
	n.pushDown(start, end)

	ans := 0
	mid := start + (end-start)/2
	if left <= mid {
		ans = max(ans, n.left.Query(start, mid, left, right))
	}
	if right > mid {
		ans = max(ans, n.right.Query(mid+1, end, left, right))
	}

	return ans
}

func (n *Node) Update(start, end, left, right, val int) {
	if left <= start && right >= end {
		n.val += val
		n.lazy += val
		fmt.Println("update:", start, end)
		return
	}
	n.pushDown(start, end)

	mid := start + (end-start)/2
	if left <= mid {
		n.left.Update(start, mid, left, right, val)
	}
	if right > mid {
		n.right.Update(mid+1, end, left, right, val)
	}
	n.pushUp(start, end)
}

func (n *Node) pushUp(start, end int) {
	n.val = max(n.left.val, n.right.val)
	if n.val > 0 {
		fmt.Println("pushup update:", start, end)
	}
}

func (n *Node) pushDown(begin, end int) {
	if n.left == nil {
		n.left = &Node{}
	}
	if n.right == nil {
		n.right = &Node{}
	}
	if n.lazy == 0 {
		return
	}

	fmt.Println("PushDown:", begin, end)
	n.left.val += n.lazy
	n.left.lazy += n.lazy
	n.right.val += n.lazy
	n.right.lazy += n.lazy

	n.lazy = 0
}

type MyCalendar struct {
	root *Node
	N    int
}

func NewMyCalendar() MyCalendar {
	return MyCalendar{
		root: &Node{},
		N:    40,
	}
}

func (c *MyCalendar) Book(startTime int, endTime int) bool {
	if c.root.Query(0, c.N, startTime, endTime-1) != 0 {
		return false
	}
	c.root.Update(0, c.N, startTime, endTime-1, 1)
	return true
}

func TestCalendar1(t *testing.T) {
	c := NewMyCalendar()
	fmt.Println(c.Book(10, 20))
	fmt.Println(c.Book(15, 25))
	fmt.Println(c.Book(21, 30))
}

type MyCalendar2 struct {
	root *Node
	N    int
}

func NewMyCalendar2() MyCalendar2 {
	return MyCalendar2{
		root: &Node{},
		N:    60,
	}
}

func (c *MyCalendar2) Book(startTime int, endTime int) bool {
	if c.root.Query(0, c.N, startTime, endTime-1) >= 2 {
		return false
	}
	c.root.Update(0, c.N, startTime, endTime-1, 1)
	return true
}

func TestCalendar2(t *testing.T) {
	c := NewMyCalendar2()
	fmt.Println(c.Book(10, 20)) // true
	fmt.Println(c.Book(50, 60)) // true
	fmt.Println(c.Book(10, 40)) // true
	fmt.Println(c.Book(5, 15))  // false
	fmt.Println(c.Book(5, 10))  // true
	fmt.Println(c.Book(25, 55)) // true
}
