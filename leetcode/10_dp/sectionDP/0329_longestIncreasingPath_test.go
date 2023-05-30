package sectionDP_test

import "math"

/*
方法：本题和1340、2713有些许相似之处
*/
func longestIncreasingPath(matrix [][]int) int {
	row, col := len(matrix), len(matrix[0])

	// 统计行、列的最大值
	rowMax := make([]int, row)
	colMax := make([]int, col)

	for i := 0; i < row; i++ {
		rowMax[i] = max(matrix[i]...)
	}
	for j := 0; j < col; j++ {
		colM := math.MinInt
		for i := 0; i < row; i++ {
			if matrix[i][j] > colM {
				colM = matrix[i][j]
			}
		}
		colMax[j] = colM
	}

	mem := make([][]int, row)
	for i := 0; i < row; i++ {
		mem[i] = make([]int, col)
		for j := 0; j < col; j++ {
			mem[i][j] = -1
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		var ret = 1
		defer func() {
			mem[i][j] = ret
		}()

		// base case
		if matrix[i][j] == rowMax[i] && matrix[i][j] == colMax[j] {
			return 1
		}

		//枚举上下左右
		for d := 0; d < 4; d++ {
			nextI := i + dx[d]
			nextJ := j + dy[d]

			if nextI >= 0 && nextI < row && nextJ >= 0 && nextJ < col && matrix[nextI][nextJ] > matrix[i][j] {
				ret = max(ret, dfs(nextI, nextJ)+1)
			}
		}
		return ret
	}

	var ans int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}
