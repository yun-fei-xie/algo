package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
931. 下降路径最小和
https://leetcode.cn/problems/minimum-falling-path-sum/
给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。

这个题是第120题的一个变形题。两个题解法的方式基本一致。
这个题可以原地dp，直接在原始数组上记录dp信息。
*/
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, 0)
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, n))
	}

	for i := 0; i < n; i++ {
		dp[n-1][i] = matrix[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			// 三个方向 找最小的那个 检查j是否越界
			var d0, d1, d2 int
			if j == 0 {
				d0 = math.MaxInt32
			} else {
				d0 = dp[i+1][j-1]
			}

			d1 = dp[i+1][j]

			if j == n-1 {
				d2 = math.MaxInt32
			} else {
				d2 = dp[i+1][j+1]
			}
			dp[i][j] = matrix[i][j] + min3(d0, d1, d2)
		}
	}
	var ans = math.MaxInt32
	for i := 0; i < n; i++ {
		if dp[0][i] < ans {
			ans = dp[0][i]
		}
	}

	return ans
}

/*
优化空间结构，直接在matrix上进行地推
*/
func minFallingPathSum2(matrix [][]int) int {
	n := len(matrix)
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			// 三个方向 找最小的那个 检查j是否越界
			var d0, d1, d2 int
			if j == 0 {
				d0 = math.MaxInt32
			} else {
				d0 = matrix[i+1][j-1]
			}

			d1 = matrix[i+1][j]

			if j == n-1 {
				d2 = math.MaxInt32
			} else {
				d2 = matrix[i+1][j+1]
			}
			matrix[i][j] = matrix[i][j] + min3(d0, d1, d2)
		}
	}
	var ans = math.MaxInt32
	for i := 0; i < n; i++ {
		if matrix[0][i] < ans {
			ans = matrix[0][i]
		}
	}

	return ans
}

/*
递归:后序遍历
*/
func minFallingPathSum3(matrix [][]int) int {
	n := len(matrix)
	var dfs func(x, y int) int // 从x,y到达底部的下降路径最小和 x和y需要进行校验 不传入越界参数
	dfs = func(x, y int) int {
		if x == n-1 {
			return matrix[x][y]
		}

		var d1, d2, d3 int
		if y-1 < 0 {
			d1 = math.MaxInt32
		} else {
			d1 = dfs(x+1, y-1)
		}

		d2 = dfs(x+1, y)
		if y+1 >= n {
			d3 = math.MaxInt32
		} else {
			d3 = dfs(x+1, y+1)
		}
		return min3(d1, d2, d3) + matrix[x][y]
	}
	var ans int = math.MaxInt32
	for i := 0; i < n; i++ {
		path := dfs(0, i)
		if path < ans {
			ans = path
		}
	}
	return ans
}

func min3(i, j, k int) int {
	var temp int
	if i > j {
		temp = j
	} else {
		temp = i
	}
	if k > temp {
		return temp
	}
	return k
}

func TestMinFallingPathSum(t *testing.T) {

	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
	fmt.Println(minFallingPathSum2([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}
