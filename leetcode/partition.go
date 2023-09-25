package leetcode

func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	p1, p2 := dummy1, dummy2
	p := head

	for p != nil {
		if p.Val < x {
			p1.Next = p
			p = p.Next
			p1 = p1.Next
			p1.Next = nil
		} else {
			p2.Next = p
			p = p.Next
			p2 = p2.Next
			p2.Next = nil
		}
	}

	p1.Next = dummy2.Next
	return dummy1.Next
}
