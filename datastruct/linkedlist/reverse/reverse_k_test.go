package reverse

func reverseKGroup(head *ListNode, k int) *ListNode {

	reverse := func (head *ListNode) *ListNode {
		var prev *ListNode
		cur, next := head, head.Next
		for cur != nil {
			cur.Next = prev
			prev = cur
			cur = next
			if next != nil {
				next = next.Next
			}
		}
		return cur
	}


	dummy := &ListNode{Next: head}
	prev, end := dummy, dummy

	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}

		next := end.Next
		end.Next = nil
		start := prev.Next
		prev.Next = reverse(start)
		start.Next = next
		
		prev = start
		end = start
	}

	return dummy.Next
}