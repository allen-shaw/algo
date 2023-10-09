package linklist

// 定义：输入以 head 开头的单链表，将这个单链表中的每两个元素翻转，
// 返回翻转后的链表头结点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next
	others := head.Next.Next

	second.Next = first
	first.Next = swapPairs(others)

	return second
}
