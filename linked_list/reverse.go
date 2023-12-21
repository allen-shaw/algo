package linkedlist

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	lp, rp := dummy, dummy
	i := 0

	for i <= right {
		if i < left-1 {
			lp = lp.Next
		}
		if i <= right {
			rp = rp.Next
		}
		i++
	}

	lp.Next = reverse(lp.Next, right-left+1)
	for lp.Next != nil {
		lp = lp.Next
	}
	lp.Next = rp

	return dummy.Next
}

// return newHead, newTail
func reverse(head *ListNode, len int) *ListNode {
	if len == 1 {
		head.Next = nil
		return head
	}

	newHead := reverse(head.Next, len-1)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right-left+1)
	}

	head.Next = reverseBetween2(head.Next, left-1, right-1)
	return head
}

var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}

	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor

	return last
}
