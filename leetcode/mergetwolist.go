package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	dummy := &ListNode{}
	p := dummy

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p = p.Next
			p1 = p1.Next
		} else {
			p.Next = p2
			p = p.Next
			p2 = p2.Next
		}
	}

	if p1 == nil {
		for p2 != nil {
			p.Next = p2
			p = p.Next
			p2 = p2.Next
		}
	}
	if p2 == nil {
		for p1 != nil {
			p.Next = p1
			p = p.Next
			p1 = p1.Next
		}
	}

	return dummy.Next
}
