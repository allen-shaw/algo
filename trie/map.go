package trie

type Map[V any] interface {
	/***** 增/改 *****/
	// 在 Map 中添加 key
	Put(key string, val V)

	/***** 删 *****/
	// 删除键 key 以及对应的值
	Remove(key string)

	/***** 查 *****/
	// 搜索 key 对应的值，不存在则返回 null
	// get("the") -> 4
	// get("tha") -> null
	Get(key string) V

	// 判断 key 是否存在在 Map 中
	// containsKey("tea") -> false
	// containsKey("team") -> true
	ContainKey(key string) bool

	// 在 Map 的所有键中搜索 query 的最短前缀
	// shortestPrefixOf("themxyz") -> "the"
	ShortestPrefixOf(query string) string

	// 在 Map 的所有键中搜索 query 的最长前缀
	// longestPrefixOf("themxyz") -> "them"
	LongestPrefixOf(query string) string

	// 搜索所有前缀为 prefix 的键
	// keysWithPrefix("th") -> ["that", "the", "them"]
	KeysWithPrefix(prefix string) []string

	// 判断是和否存在前缀为 prefix 的键
	// hasKeyWithPrefix("tha") -> true
	// hasKeyWithPrefix("apple") -> false
	HasKeyWithPrefix(prefix string) bool

	// 通配符 . 匹配任意字符，搜索所有匹配的键
	// keysWithPattern("t.a.") -> ["team", "that"]
	KeysWithPattern(pattern string) []string

	// 通配符 . 匹配任意字符，判断是否存在匹配的键
	// hasKeyWithPattern(".ip") -> true
	// hasKeyWithPattern(".i") -> false
	HasKeyWithPattern(pattern string) bool

	// 返回 Map 中键值对的数量
	Size() int
}

func NewMap[V any]() Map[V] {
	return newTripMap[V]()
}

const (
	R = 256
)

type tripNode[V any] struct {
	val      V
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
func (*tripMap[V]) ContainKey(key string) bool {
	panic("unimplemented")
}

// Get implements Map.
func (m *tripMap[V]) Get(key string) V {
	n := m.getNode(m.root, key)
	var v V
	if n == nil {
		return v
	}
	return n.val
}

// HasKeyWithPattern implements Map.
func (*tripMap[V]) HasKeyWithPattern(pattern string) bool {
	panic("unimplemented")
}

// HasKeyWithPrefix implements Map.
func (*tripMap[V]) HasKeyWithPrefix(prefix string) bool {
	panic("unimplemented")
}

// KeysWithPattern implements Map.
func (*tripMap[V]) KeysWithPattern(pattern string) []string {
	panic("unimplemented")
}

// KeysWithPrefix implements Map.
func (*tripMap[V]) KeysWithPrefix(prefix string) []string {
	panic("unimplemented")
}

// LongestPrefixOf implements Map.
func (*tripMap[V]) LongestPrefixOf(query string) string {
	panic("unimplemented")
}

// Put implements Map.
func (*tripMap[V]) Put(key string, val V) {
	panic("unimplemented")
}

// Remove implements Map.
func (*tripMap[V]) Remove(key string) {
	panic("unimplemented")
}

// ShortestPrefixOf implements Map.
func (*tripMap[V]) ShortestPrefixOf(query string) string {
	panic("unimplemented")
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
