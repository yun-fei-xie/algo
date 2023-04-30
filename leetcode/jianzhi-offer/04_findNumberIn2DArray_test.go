package jianzhi_offer

import (
	"fmt"
	"testing"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {

	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	var dfs func(x, y int, target int) bool
	dfs = func(x, y int, target int) bool {
		if outOfBound(m, n, x, y) {
			return false
		}
		if matrix[x][y] == target {
			return true
		} else if target > matrix[x][y] {
			return dfs(x+1, y, target)
		} else {
			return dfs(x, y-1, target)
		}
	}

	return dfs(0, n-1, target)
}

func outOfBound(m, n, x, y int) bool {
	if x < 0 || x >= m || y < 0 || y >= n {
		return true
	}
	return false
}

func TestFindNumberIn2DArray(t *testing.T) {
	// 从右上角看这个二维数组，就是一颗BST
	fmt.Println(findNumberIn2DArray([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 100))
}
