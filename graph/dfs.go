package graph

import (
	"fmt"
)

var dir = []int{-1, 0, 1, 0, -1}

func areaRound(board [][]byte) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	fmt.Println(board)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	board[i][j] = '#'
	for k := 0; k < 4; k++ {
		x, y := i+dir[k], j+dir[k+1]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] == 'O' {
			dfs(board, x, y)
		}
	}
}

var (
	eans     bool
	evisited [][]bool
	epath    []byte
)

func exist(board [][]byte, word string) bool {
	eans = false
	evisited = make([][]bool, len(board))
	for i := 0; i < len(evisited); i++ {
		evisited[i] = make([]bool, len(board[0]))
	}
	epath = make([]byte, 0, len(word))

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !evisited[i][j] {
				edfs(board, i, j, word, 0)
				if eans {
					return true
				}
			}
		}
	}
	return false
}

func edfs(board [][]byte, i, j int, word string, k int) {
	if eans || k >= len(word) {
		return
	}

	if board[i][j] != word[k] {
		return
	}

	evisited[i][j] = true
	epath = append(epath, board[i][j])
	if len(epath) == len(word) {
		eans = true
		return
	}

	for idx := 0; idx < 4; idx++ {
		x, y := i+dir[idx], j+dir[idx+1]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && !evisited[x][y] && k+1 < len(word) {
			edfs(board, x, y, word, k+1)
		}
	}

	epath = epath[:len(epath)-1]
	evisited[i][j] = false
}

var (
	cccvisited []bool
)

func countCompleteComponents(n int, edges [][]int) int {
	g := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		a, b := edge[0], edge[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	ans := 0
	cccvisited = make([]bool, n)

	for v := range g {
		if !cccvisited[v] {
			cccdfs(g, v)
			ans++
		}
	}

	return ans
}

func cccdfs(g [][]int, v int) {
	if cccvisited[v] {
		return
	}

	cccvisited[v] = true
	edges := g[v]
	for _, n := range edges {
		cccdfs(g, n)
	}
}

var (
	fcnvisited []bool
)

func findCircleNum(isConnected [][]int) int {
	fcnvisited = make([]bool, len(isConnected))
	ans := 0

	for i := range isConnected {
		if !fcnvisited[i] {
			fcndfs(isConnected, i)
			ans++
		}
	}

	return ans
}

func fcndfs(g [][]int, i int) {
	if fcnvisited[i] {
		return
	}

	fcnvisited[i] = true
	edges := g[i]
	for v, c := range edges {
		if c == 0 {
			continue
		}
		fcndfs(g, v)
	}
}
