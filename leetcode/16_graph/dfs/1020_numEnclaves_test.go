package dfs__test

import (
	"fmt"
	"testing"
)

/*
1020. 飞地的数量
https://leetcode.cn/problems/number-of-enclaves/description/

给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。

方法：在每一个边界上进行dfs,然后打上标记。最后没有标记的位置并且grid[x][y]==1 的位置就是飞地的格子。
	然后再进行一次遍历，统计飞地的格子总数量即可。
*/

func numEnclaves(grid [][]int) int {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	row, col := len(grid), len(grid[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		visited[i][j] = true
		for k := 0; k < 4; k++ {
			nextI := i + dx[k]
			nextJ := j + dy[k]
			if nextI >= 0 && nextI < row && nextJ >= 0 && nextJ < col {
				if !visited[nextI][nextJ] && grid[nextI][nextJ] == 1 {
					dfs(nextI, nextJ)
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		if !visited[i][0] && grid[i][0] == 1 {
			dfs(i, 0)
		}

		if !visited[i][col-1] && grid[i][col-1] == 1 {
			dfs(i, col-1)
		}
	}

	for j := 0; j < col; j++ {
		if !visited[0][j] && grid[0][j] == 1 {
			dfs(0, j)
		}
		if !visited[row-1][j] && grid[row-1][j] == 1 {
			dfs(row-1, j)
		}
	}

	var ans int
	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func TestNumEnclaves(t *testing.T) {
	fmt.Println(numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
}
