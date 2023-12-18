package doublepointer

func deleteDuplicates2(head *ListNode) *ListNode {
	dammy := &ListNode{Next: head}

	prev, slow, fast := dammy, dammy.Next, dammy.Next
	haveDup := false

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		if fast.Val != slow.Val {
			if haveDup {
				prev.Next = fast
				slow = fast
			} else {
				prev = slow
				slow = slow.Next
			}
		} else {
			haveDup = true
		}
	}

	if haveDup {
		prev.Next = fast
	}

	return dammy.Next
}

func deleteDuplicates2R(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates2R(head.Next)
		return head
	} else {
		move := head
		for move.Val == head.Val {
			move = move.Next
		}
		return deleteDuplicates2R(move)
	}
}

func deleteDuplicatesP(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev, cur := dummy, head

	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}

		if prev.Next == cur {
			prev = prev.Next
		} else {
			prev.Next = cur.Next
		}
		cur = cur.Next
	}

	return dummy.Next
}
