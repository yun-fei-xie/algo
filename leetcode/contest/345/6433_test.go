package _45

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/contest/weekly-contest-345/problems/maximum-number-of-moves-in-a-grid/

6433. 矩阵中移动的最大次数

给你一个下标从 0 开始、大小为 m x n 的矩阵 grid ，矩阵由若干 正 整数组成。
你可以从矩阵第一列中的 任一 单元格出发，按以下方式遍历 grid ：
从单元格 (row, col) 可以移动到 (row - 1, col + 1)、(row, col + 1) 和 (row + 1, col + 1) 三个单元格中任一满足值 严格 大于当前单元格的单元格。
返回你在矩阵中能够 移动 的 最大 次数。

输入：grid = [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]
输出：3
解释：可以从单元格 (0, 0) 开始并且按下面的路径移动：
- (0, 0) -> (0, 1).
- (0, 1) -> (1, 2).
- (1, 2) -> (2, 3).
可以证明这是能够移动的最大次数。

方法：
仔细观察可以发现，每一列只能走一个格子。所有，移动的最大次数其实就是路径的最大长度。
*/
func maxMoves(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	mem := make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = -1
		}
	}
	var ans = 0
	// 传入坐标（x,y）返回从x,y开始移动能够移动的最大次数
	var dfs func(x int, y int) int
	dfs = func(x int, y int) int {
		if y == n-1 {
			return 0
		}
		if mem[x][y] != -1 {
			return mem[x][y]
		}

		var res int
		// 处理上下界
		for k := max(0, x-1); k <= min(m-1, x+1); k++ {
			if grid[x][y] < grid[k][y+1] {
				res = max(res, dfs(k, y+1)+1)
			}
		}
		mem[x][y] = res
		return mem[x][y]
	}

	for i := 0; i < m; i++ {
		ans = max(ans, dfs(i, 0))
	}
	return ans
}

func max(num ...int) int {
	m := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > m {
			m = num[i]
		}
	}
	return m
}

func min(num ...int) int {
	m := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] < m {
			m = num[i]
		}
	}
	return m
}

func TestMaxMoves(t *testing.T) {
	//fmt.Println(maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}))
	fmt.Println(maxMoves([][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}))
}
