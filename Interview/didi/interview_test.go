package didi

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

// 你应当 保留
type ListNode struct {
	val  int
	next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessDummy, largeDummy := &ListNode{}, &ListNode{}
	lessCur, largeCur := lessDummy, largeDummy

	for cur := head; cur != nil; {
		if cur.val < x {
			lessCur.next = cur
			lessCur = lessCur.next
		} else {
			largeCur.next = cur
			largeCur = largeCur.next
		}
		next := cur.next 
		cur.next = nil
		cur = next
	}

	lessCur.next = largeDummy.next
	return lessDummy.next
}
