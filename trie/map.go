package trie

type Trip[V any] interface {
	/***** 增/改 *****/
	// 在 Map 中添加 key
	Put(key string, val V)

	/***** 删 *****/
	// 删除键 key 以及对应的值
	Remove(key string)

	/***** 查 *****/
	// 搜索 key 对应的值，不存在则返回 v, false
	// Get("the") -> 4, true
	// Get("tha") -> 0, false
	Get(key string) (V, bool)

	// 判断 key 是否存在在 Map 中
	// ContainKey("tea") -> false
	// ContainKey("team") -> true
	ContainKey(key string) bool

	// 在 Map 的所有键中搜索 query 的最短前缀
	// ShortestPrefixOf("themxyz") -> "the"
	ShortestPrefixOf(query string) string

	// 在 Map 的所有键中搜索 query 的最长前缀
	// LongestPrefixOf("themxyz") -> "them"
	LongestPrefixOf(query string) string

	// 搜索所有前缀为 prefix 的键
	// KeysWithPrefix("th") -> ["that", "the", "them"]
	KeysWithPrefix(prefix string) []string

	// 判断是和否存在前缀为 prefix 的键
	// HasKeyWithPrefix("tha") -> true
	// HasKeyWithPrefix("apple") -> false
	HasKeyWithPrefix(prefix string) bool

	// 通配符 . 匹配任意字符，搜索所有匹配的键
	// KeysWithPattern("t.a.") -> ["team", "that"]
	KeysWithPattern(pattern string) []string

	// 通配符 . 匹配任意字符，判断是否存在匹配的键
	// HasKeyWithPattern(".ip") -> true
	// HasKeyWithPattern(".i") -> false
	HasKeyWithPattern(pattern string) bool

	// 返回 Map 中键值对的数量
	Size() int
}

func NewTrip[V any]() Trip[V] {
	return newTripMap[V]()
}

const (
	R = 256
)

type tripNode[V any] struct {
	val      V
	ok       bool
	children []*tripNode[V]
}

func newTripNode[V any]() *tripNode[V] {
	n := &tripNode[V]{}
	n.children = make([]*tripNode[V], R)
	return n
}

type tripMap[V any] struct {
	size int
	root *tripNode[V]
}

func newTripMap[V any]() *tripMap[V] {
	tm := &tripMap[V]{}
	return tm
}

// ContainKey implements Map.
func (m *tripMap[V]) ContainKey(key string) bool {
	_, ok := m.Get(key)
	return ok
}

// Get implements Map.
func (m *tripMap[V]) Get(key string) (V, bool) {
	n := m.getNode(m.root, key)
	var v V
	if n == nil || !n.ok {
		return v, false
	}
	return n.val, true
}

// HasKeyWithPattern implements Map.
func (m *tripMap[V]) HasKeyWithPattern(pattern string) bool {
	return m.hasKeyWithPattern(m.root, pattern, 0)
}

func (m *tripMap[V]) hasKeyWithPattern(root *tripNode[V], pattern string, idx int) bool {
	if root == nil {
		return false
	}

	if idx == len(pattern) {
		return root.ok
	}

	c := pattern[idx]
	if c == '.' {
		for i := range m.root.children {
			if m.hasKeyWithPattern(root.children[i], pattern, idx+1) {
				return true
			}
		}
		return false
	} else {
		return m.hasKeyWithPattern(root.children[c], pattern, idx+1)
	}
}

// HasKeyWithPrefix implements Map.
func (m *tripMap[V]) HasKeyWithPrefix(prefix string) bool {
	n := m.getNode(m.root, prefix)
	return n == nil
}

// KeysWithPattern implements Map.
func (m *tripMap[V]) KeysWithPattern(pattern string) []string {
	keys := make([]string, 0)
	path := make([]byte, 0)
	m.traverse2(m.root, pattern, 0, path, &keys)
	return keys
}

func (m *tripMap[V]) traverse2(root *tripNode[V], pattern string, i int, path []byte, keys *[]string) {
	if root == nil {
		return
	}

	if i == len(pattern) {
		if root.ok {
			*keys = append(*keys, string(path))
		}
		return
	}

	c := pattern[i]
	if c == '.' {
		for j := 0; j < len(root.children); j++ {
			path = append(path, byte(j))
			m.traverse2(root.children[j], pattern, i+1, path, keys)
			path = path[:len(path)-1]
		}
	} else {
		m.traverse2(root.children[c], pattern, i+1, append(path, byte(c)), keys)
	}
}

func (m *tripMap[V]) KeysWithPrefix(prefix string) []string {
	node := m.getNode(m.root, prefix)
	if node == nil {
		return nil
	}

	keys := make([]string, 0)
	m.traverse(node, []byte(prefix), &keys)
	return keys
}

func (m *tripMap[V]) traverse(root *tripNode[V], path []byte, keys *[]string) {
	if root == nil {
		return
	}

	if root.ok {
		*keys = append(*keys, string(path))
	}

	for i := range root.children {
		path = append(path, byte(i))
		m.traverse(root.children[i], path, keys)
		path = path[:len(path)-1]
	}
}

// LongestPrefixOf implements Map.
func (m *tripMap[V]) LongestPrefixOf(query string) string {
	p := m.root
	maxSize := 0
	for i := 0; i < len(query); i++ {
		if p == nil {
			break
		}

		if p.ok {
			maxSize = i
		}

		c := query[i]
		p = p.children[c]
	}

	if p != nil && p.ok {
		return query
	}

	return query[:maxSize]
}

// Put implements Map.
func (m *tripMap[V]) Put(key string, val V) {
	if !m.ContainKey(key) {
		m.size++
	}

	m.root = m.put(m.root, key, 0, val)
}

func (m *tripMap[V]) put(root *tripNode[V], key string, i int, v V) *tripNode[V] {
	if root == nil {
		root = newTripNode[V]()
	}

	if i == len(key) {
		root.val = v
		root.ok = true
		return root
	}

	c := key[i]
	root.children[c] = m.put(root.children[c], key, i+1, v)

	return root
}

// Remove implements Map.
func (m *tripMap[V]) Remove(key string) {
	if !m.ContainKey(key) {
		return
	}

	m.root = m.remove(m.root, key, 0)
}

func (m *tripMap[V]) remove(root *tripNode[V], key string, i int) *tripNode[V] {
	if root == nil {
		return nil
	}

	if i == len(key) {
		root.ok = false
	} else {
		c := key[i]
		root.children[c] = m.remove(root.children[c], key, i+1)
	}

	if root.ok {
		return root
	}

	for _, n := range root.children {
		if n.ok {
			return root
		}
	}

	return nil
}

// ShortestPrefixOf implements Map.
func (m *tripMap[V]) ShortestPrefixOf(query string) string {
	p := m.root
	for i := 0; i < len(query); i++ {
		if p == nil {
			return ""
		}

		if p.ok {
			return query[:i]
		}

		c := query[i]
		p = p.children[c]
	}

	if p != nil && p.ok {
		return query
	}

	return "s"
}

// Size implements Map.
func (m *tripMap[V]) Size() int {
	return m.size
}

func (m *tripMap[V]) getNode(node *tripNode[V], key string) *tripNode[V] {
	p := node
	for i := 0; i < len(key); i++ {
		if p == nil {
			return nil
		}

		c := key[i]
		p = p.children[c]
	}

	return p
}
