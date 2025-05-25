package trie

const R = 26

var emptyByte byte

type Node struct {
	val      byte
	children []*Node
}

func newNode(val byte) *Node {
	n := &Node{}
	n.val = val
	n.children = make([]*Node, R)
	return n
}

type Trie struct {
	root *Node
}

func (t *Trie) getNode(root *Node, key string) *Node {
	p := root
	for _, b := range key {
		if p == nil {
			return nil
		}
		p = p.children[b-'a']
	}
	return p
}

func (t *Trie) keysWithPrefix(prefix string) []string {
	root := t.getNode(t.root, prefix)
	if root == nil {
		return nil
	}

	res := make([]string, 0)
	dfs(root, make([]byte, 0), res)
	return res
}

func dfs(root *Node, str []byte, res []string) {
	if root == nil {
		return
	}

	if root.val != emptyByte {
		res = append(res, string(str))
	}

	for _, child := range root.children {
		str = append(str, child.val)
		dfs(child, str, res)
		str = str[:len(str)-1]
	}
}
