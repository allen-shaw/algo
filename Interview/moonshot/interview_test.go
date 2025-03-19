package moonshot

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

type Request struct {
	timestamp    int64
	inputLength  int
	outputLength int
	hashIds      []int
}

type Block struct {
	id int
}

type Cache[T any] interface {
	Get(key int) (T, bool)
	Set(key int, value T)
	Evict() T
}

type Node struct {
	id     int
	blocks Cache[*Block]
}

func NewNode(id int, M int) *Node {
	n := &Node{id: id}
	n.blocks = NewLRUCache[*Block](M)
	return n
}

func (n *Node) getBlock(hashId int) (*Block, bool) {
	return n.blocks.Get(hashId)
}

type Metrics struct {
	reqs int
	hit  int
}

func (m *Metrics) Hit() {
	m.hit++
}

func (m *Metrics) Request() {
	m.reqs++
}

func (m *Metrics) HitRate() float64 {
	return float64(m.hit) / float64(m.reqs)
}

type Coordinator struct {
	nodes    []*Node
	selector Selector
	metrics  Metrics
}

func NewCoordinator(N, M int) *Coordinator {
	nodes := make([]*Node, N)
	for i := 0; i < N; i++ {
		nodes[i] = NewNode(i, M)
	}
	return &Coordinator{
		nodes:    nodes,
		selector: NewRandomSelector(nodes),
	}
}

type RandomSelector struct {
	router map[int][]*Node
	nodes  []*Node
}

func NewRandomSelector(nodes []*Node) *RandomSelector {
	return &RandomSelector{
		router: make(map[int][]*Node),
		nodes:  nodes,
	}
}

func (s *RandomSelector) Select(hashId int) (*Node, error) {
	// TODO: check len(s.nodes) > 0
	if nodes, ok := s.router[hashId]; ok {
		idx := rand.IntN(len(nodes))
		return nodes[idx], nil
	}
	idx := rand.IntN(len(s.nodes))
	return s.nodes[idx], nil
}

func (s *RandomSelector) AddRouter(hashId int, node *Node) {
	s.router[hashId] = append(s.router[hashId], node)
}

func (c *Coordinator) SelectNode(req Request) (*Node, error) {
	nodesMap := make(map[*Node]int)
	for _, hashId := range req.hashIds {
		c.metrics.Request()
		node, err := c.selector.Select(hashId)
		if err != nil {
			return nil, err
		}
		_, ok := node.getBlock(hashId)
		if ok {
			nodesMap[node]++
			c.metrics.Hit()
		}
		node.blocks.Set(hashId, &Block{})
		c.selector.AddRouter(hashId, node)
	}

	return maxNodes(nodesMap), nil
}

func maxNodes(m map[*Node]int) *Node {
	var node *Node
	var maxHit int

	for n, hit := range m {
		if hit > maxHit {
			node = n
			maxHit = hit
		}
	}
	return node
}

type Selector interface {
	Select(hashId int) (*Node, error)
	AddRouter(hashId int, node *Node)
}

func simulate() {
	N, M := 3, 5
	c := 10
	requests := make([]Request, c)
	for i := 0; i < c; i++ {
		requests[i] = Request{
			timestamp:    int64(i),
			inputLength:  1,
			outputLength: 1,
			hashIds:      []int{1},
		}
	}

	co := NewCoordinator(N, M)

	for _, req := range requests {
		co.SelectNode(req)
	}

	fmt.Println(co.metrics.HitRate())
}

func Test_simulate(t *testing.T) {
	simulate()
}
