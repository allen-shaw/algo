package trie

type Node struct {
	val      string
	children []*Node
}

type TrieMap struct {
	root *Node
}

func New() *TrieMap {
	return &TrieMap{root: &Node{}}
}

func (m *TrieMap) Put(key, val string) {
	m.root = m.put(m.root, key, val, 0)
}

func (m *TrieMap) put(node *Node, key, val string, index int) *Node {
	if node == nil {
		node = &Node{children: make([]*Node, 26)}
	}
	if len(key) == index {
		node.val = val
		return node
	}

	c := int(key[index] - 'a')
	node.children[c] = m.put(node.children[c], key, val, index+1)
	return node
}

func (m *TrieMap) Get(key string) string {
	n := m.getNode(m.root, key)
	if n == nil || n.val == "" {
		return ""
	}
	return n.val
}

func (m *TrieMap) getNode(node *Node, key string) *Node {
	p := node
	for i := 0; i < len(key); i++ {
		if p == nil {
			return nil
		}

		p = p.children[key[i]]
	}
	return p
}

func (m *TrieMap) keysWithPrefix(prefix string) []string {
	n := m.getNode(m.root, prefix)
	if n == nil {
		return nil
	}

	var keys []string
	var dfs func(node *Node, path []byte)
	dfs = func(node *Node, path []byte) {
		if node == nil {
			return
		}

		if node.val != "" {
			keys = append(keys, string(path))
		}

		for i := 0; i < len(n.children); i++ {
			path = append(path, byte(i+'a'))
			dfs(node.children[i], path)
			path = path[len(path)-1:]
		}

	}

	dfs(n, []byte(prefix))
	return keys
}
