package graph

import "fmt"

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
