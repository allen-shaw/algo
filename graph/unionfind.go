package graph

type UnionFind struct {
	// 连通分量的数量
	count int
	// 父节点，节点 x 的父节点是 parent[x]
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		count:  n,
		parent: make([]int, n),
	}

	for i := range uf.parent {
		uf.parent[i] = i
	}

	return uf
}

// 将 p 和 q 连接
func (uf *UnionFind) Union(p, q int) {
	rootP, rootQ := uf.find(p), uf.find(q)
	if rootP == rootQ {
		return
	}

	// 将rootP挂到rootQ上，rootP的父节点为rootQ
	uf.parent[rootP] = rootQ
	uf.count--
}

// 判断 p 和 q 是否连通
func (uf *UnionFind) Connect(p, q int) bool {
	rootP, rootQ := uf.find(p), uf.find(q)
	return rootP == rootQ
}

// 返回图中有多少个连通分量
func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x { // 当找到根节点时，parent[x] == x
		uf.parent[x] = uf.find(uf.parent[x]) // 逐级往上找，直到根节点, 然后将x的parent，parent的parent都连到root
	}

	return uf.parent[x]
}

func countComponents(n int, edges [][]int) int {
	uf := NewUnionFind(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}

	return uf.Count()
}

// type Cell struct {
// 	v     int
// 	equal bool
// }

// func equationsPossible(equations []string) bool {
// 	g := buildGraphWithWeight(equations)
// 	fmt.Println(g)

// 	visited := make([]bool, len(g))
// 	color := make([]int, len(g))

// 	q := make([]int, 0, len(g))
// 	q = append(q, 0)
// 	visited[0] = true
// 	color[0] = 0

// 	for len(q) > 0 {
// 		v := q[0]
// 		q = q[1:]

// 		for _, cell := range g[v] {
// 			if visited[cell.v] {
// 				// 已经访问过，查看是否符合
// 				if (cell.equal && color[cell.v] != color[v]) ||
// 					(!cell.equal && color[cell.v] == color[v]) {
// 					fmt.Println("node:", v, cell, color)
// 					return false
// 				}
// 			} else {
// 				visited[cell.v] = true
// 				if cell.equal {
// 					color[cell.v] = color[v]
// 				} else {
// 					color[cell.v] = cell.v
// 				}
// 				q = append(q, cell.v)
// 			}
// 		}
// 	}

// 	return true
// }

// func buildGraphWithWeight(equations []string) [][]Cell {
// 	g := make([][]Cell, 26)
// 	for _, equation := range equations {
// 		v1, v2 := int(equation[0]-'a'), int(equation[3]-'a')
// 		equal := equation[1] == '='
// 		g[v1] = append(g[v1], Cell{v: v2, equal: equal})
// 		g[v2] = append(g[v2], Cell{v: v1, equal: equal})
// 	}

// 	return g
// }

func equationsPossible(equations []string) bool {
	uf := NewUnionFind(26)

	for _, equation := range equations {
		if equation[1] == '=' {
			v1, v2 := int(equation[0]-'a'), int(equation[3]-'a')
			uf.Union(v1, v2)
		}
	}

	for _, equation := range equations {
		if equation[1] == '!' {
			v1, v2 := int(equation[0]-'a'), int(equation[3]-'a')
			if uf.Connect(v1, v2) {
				return false
			}
		}
	}

	return true
}
