package linkedlist

var ln *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	ln = head
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}

	return buildBST(0, n-1)

}

func buildBST(l, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := (l + r) / 2
	left := buildBST(l, mid-1)

	node := &TreeNode{Val: ln.Val}
	ln = ln.Next

	right := buildBST(mid+1, r)
	node.Left = left
	node.Right = right

	return node
}
