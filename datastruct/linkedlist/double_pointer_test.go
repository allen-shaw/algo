package linkedlist

import (
	"fmt"
	"testing"
)

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for fast != nil {
		if fast.Next != nil && fast.Next.Val == fast.Val {
			for fast.Next != nil && fast.Val == fast.Next.Val {
				fast = fast.Next
			}
			fast = fast.Next
		} else {
			slow.Next = fast
			slow = slow.Next
			fast = fast.Next
		}
	}
	slow.Next = fast
	return dummy.Next
}

// d, 1, 2, 3, 3, 4, 4, 5
//    	 s        f

func Test_deleteDuplicates(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 4, 5}
	head := toLinkedList(nums)
	out := deleteDuplicates(head)
	fmt.Println(toArray(out))
}

func deleteDuplicatesUnsorted2(head *ListNode) *ListNode {
	m := make(map[int]int)
	for p := head; p != nil; p = p.Next {
		m[p.Val]++
	}

	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for fast != nil {
		for fast != nil && m[fast.Val] > 1 {
			fast = fast.Next
		}
		if fast != nil {
			slow.Next = fast
			slow = slow.Next
			fast = fast.Next
		}
	}
	slow.Next = fast

	return dummy.Next
}

// d,2,1,1,2
//
//	s     f
func Test_deleteDuplicatesUnsorted2(t *testing.T) {
	head1 := toLinkedList([]int{1, 2, 3, 2})
	head2 := toLinkedList([]int{2, 1, 1, 2})
	head3 := toLinkedList([]int{3, 2, 2, 1, 3, 2, 4})

	fmt.Println(toArray(deleteDuplicatesUnsorted2(head1)))
	fmt.Println(toArray(deleteDuplicatesUnsorted2(head2)))
	fmt.Println(toArray(deleteDuplicatesUnsorted2(head3)))
}

func nthUglyNumber(n int) int {
	ans := make([]int, n)
	nums2, nums3, nums5 := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	nums2[0] = 1
	nums3[0] = 1
	nums5[0] = 1
	p2, p3, p5 := 0, 0, 0

	for i := 0; i < n; i++ {
		val := min(nums2[p2], nums3[p3], nums5[p5])
		ans[i] = val
		if val == nums2[p2] {
			p2++
		}
		nums2[p2] = val * 2

		if val == nums3[p3] {
			p3++
			nums3[p3] = val * 3
		}
		if val == nums5[p5] {
			p5++
			nums5[p5] = val * 5
		}
	}
	fmt.Println(nums2, nums3, nums5)
	fmt.Println(ans)
	return ans[n-1]
}

func Test_nthUglyNumber(t *testing.T) {
	nthUglyNumber(10)
}


