package disjointSet

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/surrounded-regions/description/?orderBy=most_votes
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

解题思路:
使用遍历的方式将边界或者与边界相连的'O'进行标记。那么剩下的'O'便是被'X'围绕的区域。
然后使用二重循环对矩阵进行遍历，将符合条件的'O'修改为'X'即可。

标记的过程可以将'O'修改为'V'，然后在二重循环的时候将'V'再修改回'O'
*/
func solve(board [][]byte) {
	row, col := len(board), len(board[0])
	// 对上下左右边界进行一次遍历
	// 左右
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][col-1] == 'O' {
			dfs(board, i, col-1)
		}
	}

	// 上下
	for i := 0; i < col; i++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
		if board[row-1][i] == 'O' {
			dfs(board, row-1, i)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'V' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

func dfs(board [][]byte, x, y int) {
	row, col := len(board), len((board)[0])
	if x < 0 || x >= row || y < 0 || y >= col { // 出错的地方在边界判断
		return
	}
	if board[x][y] == 'O' {
		board[x][y] = 'V'
		dfs(board, x+1, y)
		dfs(board, x-1, y)
		dfs(board, x, y+1)
		dfs(board, x, y-1)
	}
}

func TestSolve(t *testing.T) {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println(board)
}

func TestSolve2(t *testing.T) {
	board := [][]byte{{'O'}}
	solve(board)
	fmt.Println(board)
}
