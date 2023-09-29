package linklist

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 找到中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 左半部分翻转,翻转至slow前一个节点
	var (
		last *ListNode
		cur  = head
		next = cur.Next
	)

	for cur != slow {
		cur.Next = last
		last = cur
		cur = next
		next = next.Next
	}

	// 判断是否是回文
	left, right := last, slow
	if fast != nil {
		// 奇数长度
		right = slow.Next
	}

	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}
