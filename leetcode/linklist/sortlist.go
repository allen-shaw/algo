package linklist

func sortList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	// 找到中点
	left, right := partition(head)
	left = sortList(left)
	right = sortList(right)

	// merge
	return merge(left, right)
}

func partition(head *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	left, right := head, slow
	prev.Next = nil

	return left, right

}

func merge(left, right *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	curLeft, curRight := left, right

	for curLeft != nil && curRight != nil {
		if curLeft.Val <= curRight.Val {
			cur.Next = curLeft
			curLeft = curLeft.Next
		} else {
			cur.Next = curRight
			curRight = curRight.Next
		}
		cur = cur.Next
	}

	for curLeft != nil {
		cur.Next = curLeft
		curLeft = curLeft.Next
		cur = cur.Next
	}
	for curRight != nil {
		cur.Next = curRight
		curRight = curRight.Next
		cur = cur.Next
	}

	return head.Next
}
