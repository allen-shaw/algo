package graph

var (
	hasCycle      bool
	canFinOnPath  map[int]struct{}
	canFinVisited []bool
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := buildGraph(numCourses, prerequisites)
	canFinOnPath = make(map[int]struct{})
	canFinVisited = make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !canFinVisited[i] {
			traverseFinish(g, i)
		}
	}
	return !hasCycle
}

func buildGraph(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, edge := range edges {
		from, to := edge[1], edge[0]
		g[from] = append(g[from], to)
	}

	return g
}

func traverseFinish(g [][]int, n int) {
	// fmt.Println()
	if _, ok := canFinOnPath[n]; ok {
		hasCycle = true
	}

	if canFinVisited[n] || hasCycle {
		return
	}

	canFinVisited[n] = true
	canFinOnPath[n] = struct{}{}
	neighborhoods := g[n]
	for _, nei := range neighborhoods {
		traverseFinish(g, nei)
	}
	delete(canFinOnPath, n)
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := buildGraph(numCourses, prerequisites)
	s := newFindOrderSolution(numCourses, g)
	s.Traverse()
	if s.hasCycle {
		return []int{}
	}
	reverse(s.res)
	return s.res
}

type findOrderSolution struct {
	num      int
	g        [][]int
	visited  []bool
	onPath   map[int]struct{}
	res      []int
	hasCycle bool
}

func newFindOrderSolution(n int, g [][]int) *findOrderSolution {
	s := &findOrderSolution{
		num:      n,
		g:        g,
		visited:  make([]bool, n),
		onPath:   make(map[int]struct{}),
		hasCycle: false,
	}

	return s
}

func (s *findOrderSolution) Traverse() {
	for i := 0; i < s.num; i++ {
		s.traverse(i)
	}
}

func (s *findOrderSolution) traverse(n int) {
	if _, ok := s.onPath[n]; ok {
		s.hasCycle = true
	}

	if s.hasCycle || s.visited[n] {
		return
	}

	s.visited[n] = true
	s.onPath[n] = struct{}{}

	neighborhoods := s.g[n]
	for _, nei := range neighborhoods {
		s.traverse(nei)
	}
	s.res = append(s.res, n)
	delete(s.onPath, n)
}

func reverse(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
