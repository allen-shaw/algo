package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1->2->3->4->5->6->7
// 1<-2<-3<-4<-5<-6<-7
// 1->2->3->4  5->6->7->
// 5->6->7->1->2->3->4
// [1], k = 1
func rotateRight(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		if fast.Next == nil {
			fast = head
		} else {
			fast = fast.Next
		}
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	last := slow
	for slow.Next != nil {
		slow = slow.Next
	}
	newHead := last
	if last.Next != nil {
		newHead = last.Next
	}
	slow.Next = head
	last.Next = nil

	return newHead
}

func Test_rotateRight(t *testing.T) {
	cases := []struct {
		nums   []int
		k      int
		expert []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
		{[]int{1}, 1, []int{1}},
	}

	for _, c := range cases {
		head := toLinkedList(c.nums)
		out := toArray(rotateRight(head, c.k))
		assert.Equal(t, c.expert, out)
	}
}
