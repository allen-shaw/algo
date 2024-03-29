package tree

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/allen-shaw/algo/tree/queue"
)

// 求一颗二叉树最大宽度
func MaxWidth(n *TreeNode) int {
	if n == nil {
		return 0
	}

	levelMap := make(map[*TreeNode]int)
	maxWidth := 0
	curLvl := 1
	curLvlCnt := 0

	q := queue.New[*TreeNode]()
	q.Enqueue(n)
	levelMap[n] = 1

	for !q.Empty() {
		n := q.Dequeue()
		lvl := levelMap[n]

		if lvl != curLvl {
			// 到下一层了，进行数据结算
			maxWidth = maxInt(curLvlCnt, maxWidth)

			// 结算后
			curLvl = lvl
			curLvlCnt = 0
		}

		curLvlCnt++
		if n.Left != nil {
			q.Enqueue(n.Left)
			levelMap[n.Left] = curLvl + 1
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
			levelMap[n.Right] = curLvl + 1
		}
	}

	// 最后一层结算
	fmt.Println(curLvl, curLvlCnt)
	maxWidth = maxInt(curLvlCnt, maxWidth)

	return maxWidth
}

func buildTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Right.Left.Right = &TreeNode{Val: 8}
	return root
}

func TestMaxWidth(t *testing.T) {
	root := buildTree2()
	n := MaxWidth(root)
	fmt.Println("max width", n)
	n2 := MaxWidth2(root)
	fmt.Println("max width", n2)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxWidth2(h *TreeNode) int {
	if h == nil {
		return 0
	}

	maxWidth := 0
	curLvlCnt := 0
	var (
		curLvlEnd  *TreeNode
		nextLvlEnd *TreeNode
	)

	q := queue.New[*TreeNode]()
	q.Enqueue(h)
	curLvlEnd = h

	for !q.Empty() {
		n := q.Dequeue()
		curLvlCnt++
		if n.Left != nil {
			q.Enqueue(n.Left)
			nextLvlEnd = n.Left
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
			nextLvlEnd = n.Right
		}

		if n == curLvlEnd {
			// 当前层到尾部,结算
			maxWidth = maxInt(curLvlCnt, maxWidth)

			curLvlCnt = 0
			curLvlEnd = nextLvlEnd
			nextLvlEnd = nil
		}
	}

	maxWidth = maxInt(curLvlCnt, maxWidth)
	return maxWidth
}

// 判断一棵二叉树是不是搜索二叉树
func IsBST(h *TreeNode) bool {
	_, _, ok := isBst(h)
	return ok
}

func isBst(n *TreeNode) (int, int, bool) {
	if n == nil {
		return 0, 0, true
	}

	lmin, lmax, lok := isBst(n.Left)
	rmin, rmax, rok := isBst(n.Right)

	ok := lok && rok && (lmax <= n.Val && n.Val <= rmin)
	min := minInts(lmin, rmin, n.Val)
	max := maxInts(lmax, rmax, n.Val)

	return min, max, ok
}

func TestIsBST(t *testing.T) {

}

// 判断一棵二叉树是不是完全二叉树

// 判断一棵二叉树是不是满二叉树

// 判断一棵二叉树是不是平衡二叉树
func IsBalanceTree(h *TreeNode) bool {
	_, ok := isBalanceTree(h)
	return ok
}

func isBalanceTree(n *TreeNode) (int, bool) {
	if n == nil {
		return 0, true
	}

	lh, lok := isBalanceTree(n.Left)
	rh, rok := isBalanceTree(n.Right)

	return maxInt(lh, rh) + 1, lok && rok && (absInt(lh-rh) <= 1)
}

func TestIsBalanceTree(t *testing.T) {

}

func absInt(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func maxInts(nums ...int) int {
	if len(nums) == 0 {
		panic("nil nums list")
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max > nums[i] {
			max = nums[i]
		}
	}
	return max
}

func minInts(nums ...int) int {
	if len(nums) == 0 {
		panic("nil nums list")
	}

	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if min < nums[i] {
			min = nums[i]
		}
	}
	return min
}

func TestForLoop(t *testing.T) {

	fn := func(i int) int {
		if i == 3 {
			return 1
		}
		return 5
	}

	for i := 0; i < fn(i); i++ {
		fmt.Println(i)
	}
}

const nilNum = -100

func TestLowestCommonAncestor(t *testing.T) {
	// root := buildATree(3, 5, 1, 6, 2, 0, 8, nilNum, nilNum, 7, 4)
	// lowestCommonAncestor1(root, )
}

func TestBuildATree(t *testing.T) {
	root := buildATree(3, 5, 1, 6, 2, 0, 8, nilNum, nilNum, 7, 4)
	//   	  3
	//  	5    1
	// 	   6 2  0 8
	//      7 4
	// 3 5 6 2 7 4 1 0 8
	PreorderTraversal(root)
}

func buildATree(nums ...int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(nums))
	for i, num := range nums {
		nodes[i] = &TreeNode{Val: num}
	}

	size := len(nodes)
	for i, node := range nodes {
		if node.Val == nilNum {
			continue
		}

		lcIdx := 2*i + 1
		rcIdx := 2*i + 2

		if lcIdx < size {
			lc := nodes[lcIdx]
			if lc.Val == nilNum {
				node.Left = nil
			} else {
				node.Left = lc
			}
		}

		if rcIdx < size {
			rc := nodes[rcIdx]
			if rc.Val == nilNum {
				node.Right = nil
			} else {
				node.Right = rc
			}
		}
	}

	return nodes[0]
}

// 求两个节点的最低公共节点
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	// 遍历求解所有的父节点
	m := make(map[*TreeNode]*TreeNode)
	set := make(map[*TreeNode]struct{})

	findParent(root, m)

	n := p
	for n != nil {
		set[n] = struct{}{}
		n = m[n]
	}

	n2 := q
	for {
		if a, ok := m[n2]; ok {
			return a
		}
		n2 = m[n2]
	}
}

func findParent(n *TreeNode, m map[*TreeNode]*TreeNode) {
	if n == nil {
		return
	}

	if n.Left != nil {
		m[n.Left] = n
		findParent(n.Left, m)
	}
	if n.Right != nil {
		m[n.Right] = n
		findParent(n.Right, m)
	}
}

func deepestLeavesSum(root *TreeNode) int {
	maxDepth := 0
	sum := 0

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, lvl int) {
		if root == nil {
			return
		}
		if lvl > maxDepth {
			maxDepth = lvl
			sum = 0
		}

		if lvl == maxDepth {
			sum += root.Val
		}

		dfs(root.Left, lvl+1)
		dfs(root.Right, lvl+1)
	}

	dfs(root, 0)
	return sum
}

func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)

	var dfs func(root *TreeNode, path *[]string)
	dfs = func(root *TreeNode, path *[]string) {
		if root == nil {
			ans = append(ans, toString(*path))
		}

		v := strconv.Itoa(root.Val)
		*path = append(*path, v)
		dfs(root.Left, path)
		dfs(root.Right, path)
		*path = (*path)[:len(*path)-1]
	}

	path := make([]string, 0)
	dfs(root, &path)
	return ans
}

func toString(nums []string) string {
	return strings.Join(nums, "->")
}

func toNums(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = ans*10 + nums[i]
	}
	return ans
}

func TestToNums(t *testing.T) {
	nums := []int{1, 2}
	ans := toNums(nums)
	fmt.Println(ans)

	a := 'a' + 2
	fmt.Println(string(a))

	// strings.
}

func longestConsecutive2(root *TreeNode) int {
	ans := 0

	var dfs func(root *TreeNode, len int, parent int)
	dfs = func(root *TreeNode, len, parent int) {
		if root == nil {
			return
		}

		if root.Val == parent+1 {
			len++
		} else {
			len = 1
		}
		ans = max(ans, len)

		dfs(root.Left, len, root.Val)
		dfs(root.Right, len, root.Val)
	}

	dfs(root, 0, root.Val)
	return ans
}

func pathSum3(nums []int) int {
	m := make(map[int]*TreeNode)
	_, _, v0 := decode(nums[0])
	root := &TreeNode{Val: v0}
	m[1] = root

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		y, x, v := decode(num)
		id := int(math.Pow(2, float64((y-1)))) + (x - 1)

		tn := &TreeNode{Val: v}
		m[id] = tn

		pid := id / 2
		p, ok := m[pid]
		if !ok {
			fmt.Println(id, pid, m, num, y, x, v)
			panic("p not exist")
		}

		if id == pid*2 {
			p.Left = tn
		} else {
			p.Right = tn
		}
	}

	travese(nil, root)

	sum := 0
	path := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		path += root.Val
		if root.Left == nil && root.Right == nil {
			sum += path
		}

		dfs(root.Left)
		dfs(root.Right)
		path -= root.Val
	}
	dfs(root)
	return sum
}

func decode(num int) (int, int, int) {
	v := num % 10
	num = num / 10
	x := num % 10
	num = num / 10
	y := num
	return y, x, v
}

func TestPathSum3(t *testing.T) {
	nums := []int{113, 215, 221}
	sum := pathSum3(nums)
	fmt.Println(sum)
}

func travese(p, root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(p, root.Val)
	travese(root, root.Left)
	travese(root, root.Right)
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := make([]*TreeNode, 0)

	for _, n := range nums {
		tn := &TreeNode{Val: n}
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			fmt.Println(tn, top)
			if tn.Val > top.Val {
				stack = stack[:len(stack)-1] // pop
				tn.Left = top
			} else { // tn.Val < top.Val
				stack = append(stack, tn)
				top.Right = tn
				break
			}
		}
		if len(stack) == 0 {
			stack = append(stack, tn)
		}
	}

	return stack[0]
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	fmt.Println(root)
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) (*TreeNode, *TreeNode)
	dfs = func(root *TreeNode) (*TreeNode, *TreeNode) {
		if root == nil {
			return nil, nil
		}

		leftMin, leftMax := dfs(root.Left)
		rightMin, rightMax := dfs(root.Right)

		root.Left = leftMax
		root.Right = rightMin

		if leftMax != nil {
			leftMax.Right = root
		}
		if rightMin != nil {
			rightMin.Left = root
		}

		return leftMin, rightMax
	}

	head, tail := dfs(root)
	tail.Right = head
	head.Left = tail
	return head
}

var memo map[int][]*TreeNode

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{}}
	}

	if v, ok := memo[n]; ok {
		return v
	}

	res := make([]*TreeNode, 0)

	i := 1
	j := n - 1 - i

	for j > 0 {
		leftTrees := allPossibleFBT(i)
		rightTrees := allPossibleFBT(j)

		for _, lt := range leftTrees {
			for _, rt := range rightTrees {
				root := &TreeNode{}
				root.Left = lt
				root.Right = rt
				res = append(res, root)
			}
		}

		i += 2
		j -= 2
	}

	memo[n] = res
	return res
}

func findLeaves(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	var dfs func(root *TreeNode) (*TreeNode, int)
	dfs = func(root *TreeNode) (*TreeNode, int) {
		if root == nil {
			return nil, 0
		}

		var ld, rd, deep int
		root.Left, ld = dfs(root.Left)
		root.Right, rd = dfs(root.Right)

		deep = max(ld, rd) + 1
		if root.Left == nil && root.Right == nil {
			for len(ans) < deep {
				ans = append(ans, make([]int, 0))
			}
			ans[deep] = append(ans[deep], root.Val)
			return nil, deep
		}

		return root, deep
	}

	dfs(root)
	return ans
}

func longestConsecutive4(root *TreeNode) int {
	maxLenght := 0

	var dfs func(root *TreeNode) (min, max int)
	dfs = func(root *TreeNode) (min, max int) {
		if root == nil {
			return 0, 0
		}

		lmin, lmax := dfs(root.Left)
		rmin, rmax := dfs(root.Right)

		min, max = 1, 1
		// 递减 求min
		if root.Left != nil && root.Val == root.Left.Val-1 {
			min = lmin + 1
		}
		if root.Right != nil && root.Val == root.Right.Val-1 {
			min = maximum(min, rmin+1)
		}

		// 递增 求max
		if root.Left != nil && root.Val == root.Left.Val+1 {
			max = lmax + 1
		}
		if root.Right != nil && root.Val == root.Right.Val+1 {
			max = maximum(max, rmax+1)
		}

		maxLenght = maximum(maxLenght, min+max-1)
		return min, max
	}

	dfs(root)

	return maxLenght
}

func maximum(nums ...int) int {
	m := nums[0]
	for _, v := range nums {
		m = max(v, m)
	}
	return m
}

func maximumAverageSubtree(root *TreeNode) float64 {
	var maxAvg float64

	var dfs func(root *TreeNode) (nodenum, sum int)
	dfs = func(root *TreeNode) (nodenum int, sum int) {
		if root == nil {
			return 0, 0
		}

		lnum, lsum := dfs(root.Left)
		rnum, rsum := dfs(root.Right)

		nodenum = lnum + rnum + 1
		sum = lsum + rsum + root.Val

		avg := float64(sum) / float64(nodenum)
		maxAvg = max(maxAvg, avg)

		return nodenum, sum
	}

	dfs(root)
	return maxAvg
}

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	tree := make([][]int, n)

	for i, p := range parents {
		if p == -1 {
			continue
		}
		tree[p] = append(tree[p], i)
	}

	maxScore := 0
	count := 0

	// 返回以i为根的子树的大小
	var dfs func(i int) int
	dfs = func(i int) int {
		score := 1
		subnums := 0
		children := tree[i]
		for _, child := range children {
			subNum := dfs(child)
			if subNum != 0 {
				score *= subNum
				subnums += subNum
			}
		}
		if (n - 1 - subnums) != 0 {
			score *= (n - 1 - subnums)
		}
		if score > maxScore {
			maxScore = score
			count = 1
		} else if score == maxScore {
			count++
		}

		return subnums + 1
	}

	dfs(0)
	return count
}

func Test_countHighestScoreNodes(t *testing.T) {
	parents := []int{-1, 2, 0} // -1, 2, 0, 2, 0}
	ans := countHighestScoreNodes(parents)
	fmt.Println(ans)
}

func treeDiameter(edges [][]int) int {
	// 最有有n+1个节点，n为边的数量
	n := len(edges)
	if n == 0 {
		return 0
	}

	inDegree := make(map[int]int)
	tree := make([][]int, n+1)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		if _, ok := inDegree[edge[0]]; !ok {
			inDegree[edge[0]] = 0
		}
		inDegree[edge[1]]++
	}

	// 找到root节点
	root := -1
	fmt.Println(inDegree)
	for n, indegree := range inDegree {
		if indegree == 0 {
			root = n
			break
		}
	}

	maxDiameter := 0

	// 返回root节点为根的最大高度
	var dfs func(root int) int
	dfs = func(root int) int {
		maxDepth, secDepth := 0, 0
		children := tree[root]
		for _, child := range children {
			depth := dfs(child)
			if depth > maxDepth {
				secDepth = maxDepth
				maxDepth = depth
			} else if depth > secDepth {
				secDepth = depth
			}
		}

		maxDiameter = max(maxDiameter, maxDepth+secDepth)
		return maxDepth + 1
	}

	dfs(root)
	return maxDiameter
}

func Test_treeDiameter(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}}
	ans := treeDiameter(edges)
	fmt.Println(ans)
}

func lowestCommonAncestor4(root *TreeNode, nodes []*TreeNode) *TreeNode {
	m := make(map[*TreeNode]struct{})
	for _, n := range nodes {
		m[n] = struct{}{}
	}

	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		if _, ok := m[root]; ok {
			return root
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		if left != nil && right != nil {
			return root
		}

		if left != nil {
			return left
		}
		return right
	}

	return dfs(root)
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var pok, qok bool

	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		if left != nil && right != nil {
			return root
		}

		if root.Val == p.Val {
			pok = true
			return root
		}
		if root.Val == q.Val {
			qok = true
			return root
		}

		if left != nil {
			return left
		}
		return right
	}

	ans := dfs(root)
	if pok && qok {
		return ans
	}
	return nil
}
