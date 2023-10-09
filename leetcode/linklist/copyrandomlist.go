package linklist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node, 0)

	for n := head; n != nil; n = n.Next {
		m[n] = &Node{Val: n.Val}
	}

	for n := head; n != nil; n = n.Next {
		m[n].Next = m[n.Next]
		m[n].Random = m[n.Random]
	}

	return m[head]
}
