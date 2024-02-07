package unionfind

import (
	"fmt"
	"testing"
)

func findRedundantConnection(edges [][]int) []int {
	// 一棵树的节点数量=边的数量+1（除了根节点没有指向父节点的边，其他每个节点都有一条指向父节点边）
	// 假设有n个节点，边数量应该是n-1，根据题意，是在一棵树上多一条边，因此边数应该是len(edge)-1; n=边数+1=len(edge)
	// 从[1,n],所以增加一个node
	uf := NewUnionFound(len(edges) + 1)

	for _, edge := range edges {
		if !uf.Union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}

type UnionFound struct {
	parent []int
	rank   []int
}

func NewUnionFound(n int) *UnionFound {
	uf := &UnionFound{
		parent: make([]int, n),
		rank:   make([]int, n),
	}

	for i := range uf.parent {
		uf.parent[i] = i
	}

	return uf
}

func (uf *UnionFound) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return false
	}

	rx, ry := uf.rank[rootX], uf.rank[rootY]
	if rx > ry {
		uf.parent[rootY] = rootX
	} else if rx < ry {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	return true
}

func (uf *UnionFound) Find(x int) int {
	if x != uf.parent[x] {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

func TestFindRedundantConnection(t *testing.T) {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	ans := findRedundantConnection(edges)
	fmt.Println(ans)
}

func TestUnionFind(t *testing.T) {
	uf := NewUnionFound(3)
	fmt.Println(uf.Find(2))
	fmt.Println(uf.Union(1, 2))
	fmt.Println(uf.parent)

	fmt.Println(uf.Find(2))

}
