package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func toLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, num := range nums {
		p.Next = &ListNode{Val: num}
		p = p.Next
	}
	return dummy.Next
}

func toArray(head *ListNode) []int {
	arr := make([]int, 0)
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}
	return arr
}

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	m := make(map[int]int)
	for p := head; p != nil; p = p.Next {
		m[p.Val]++
	}

	p, q := dummy, head
	for q != nil {
		if m[q.Val] == 1 {
			if p.Next == q {
				p = p.Next
			} else {
				p.Next = q
				p = q
			}
		}
		q = q.Next
	}

	return dummy.Next
}

func Test_deleteDuplicatesUnsorted(t *testing.T) {
	testCases := []struct {
		name   string
		input  *ListNode
		expect []int
	}{
		{"case3", toLinkedList([]int{3, 2, 2, 1, 3, 2, 4}), []int{1, 4}},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			out := deleteDuplicatesUnsorted(c.input)
			assert.Equal(t, c.expect, toArray(out))
		})
	}
}
