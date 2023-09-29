package linklist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p1, p2, p := l1, l2, dummy
	cf := 0

	for p1 != nil || p2 != nil {
		v1, v2 := 0, 0
		if p1 != nil {
			v1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			v2 = p2.Val
			p2 = p2.Next
		}

		n := v1 + v2 + cf
		if n >= 10 {
			n = n % 10
			cf = 1
		} else {
			cf = 0
		}
		p.Next = &ListNode{Val: n}
		p = p.Next
	}

	if cf == 1 {
		p.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}
