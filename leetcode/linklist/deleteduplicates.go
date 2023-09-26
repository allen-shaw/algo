package linklist

func deleteDuplicates(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}

		fast = fast.Next
	}

	slow.Next = nil
	return head
}
