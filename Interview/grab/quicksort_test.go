package grab

import (
	"fmt"
	"testing"
)

func quickSort() {

}

type ListNode struct {
	val  int
	next *ListNode
}

func quickSortLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	pivot := head.val
	lessDummy, equalDummy, largeDummy := &ListNode{}, &ListNode{}, &ListNode{}
	less, equal, large := lessDummy, equalDummy, largeDummy
	curr := head
	for curr != nil {
		if curr.val < pivot {
			less.next = curr
			less = less.next
		} else if curr.val == pivot {
			equal.next = curr
			equal = equal.next
		} else {
			large.next = curr
			large = large.next
		}

		curr = curr.next
	}
	less.next = nil
	equal.next = nil
	large.next = nil

	less = quickSortLinkedList(lessDummy.next)
	large = quickSortLinkedList(largeDummy.next)

	return concat(less, equalDummy.next, large)
}

func concat(less, equal, large *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for curr != nil {
		if less != nil {
			curr.next = less
			less = less.next
		} else if equal != nil {
			curr.next = equal
			equal = equal.next
		} else if large != nil {
			curr.next = large
			large = large.next
		}
		curr = curr.next
	}
	return dummy.next
}

func Test_quickSortLinkedList(t *testing.T) {
	head := arrayToList([]int{3, 3, 1, 2, 3})
	head = quickSortLinkedList(head)
	printLinkedList(head)
}

func arrayToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for i := range nums {
		curr.next = &ListNode{val: nums[i]}
		curr = curr.next
	}
	return dummy.next
}

func printLinkedList(head *ListNode) {
	for curr := head; curr != nil; curr = curr.next {
		if curr == head {
			fmt.Printf("%v", curr.val)
		} else {
			fmt.Printf("->%v", curr.val)
		}
	}
	fmt.Println()
}
