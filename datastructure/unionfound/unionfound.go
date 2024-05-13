package unionfound

type UnionFound struct {
	parents []int
	size    int
}

func New(size int) *UnionFound {
	uf := &UnionFound{
		parents: make([]int, size),
		size:    size,
	}
	for i := 0; i < size; i++ {
		uf.parents[i] = i
	}
	return uf
}

func (uf *UnionFound) Union(x, y int) {
	rootX, rootY := uf.root(x), uf.root(y)
	if rootX == rootY {
		return
	}
	uf.parents[rootX] = rootY
	uf.size--
}

func (uf *UnionFound) IsConnected(x, y int) bool {
	rootX, rootY := uf.root(x), uf.root(y)
	return rootX == rootY
}

func (uf *UnionFound) Count() int {
	return uf.size
}

func (uf *UnionFound) root(x int) int {
	if x != uf.parents[x] {
		uf.parents[x] = uf.root(uf.parents[x])
	}
	return uf.parents[x]
}
