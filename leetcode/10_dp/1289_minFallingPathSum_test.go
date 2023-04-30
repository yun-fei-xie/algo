package _0_dp

import "math"

/*
https://leetcode.cn/problems/minimum-falling-path-sum-ii/description/
这道题有一个需要注意的点：非零偏移下降路径 定义为：从 arr 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。
不在同一列，但是不代表在相邻列。

本题也是直接在原数组上进行dp。空间复杂度为O（1）
时间复杂度O(n^3)
*/
func minFallingPathSumII(grid [][]int) int {
	n := len(grid)

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			var mdp int = math.MaxInt32
			for k := 0; k < n; k++ { // 枚举出了当前列之外的所有列
				if k == j {
					continue
				}
				if grid[i+1][k] < mdp {
					mdp = grid[i+1][k]
				}
			}
			grid[i][j] += mdp
		}
	}
	var ans int = math.MaxInt32
	for i := 0; i < n; i++ {
		if ans > grid[0][i] {
			ans = grid[0][i]
		}
	}
	return ans
}

/*
优化：
在向上递推的过程中，求解第i行的状态时，在进行状态转移的时候，都会枚举i-1行的最小值。
这个最小值不能和当前列相同。因此，需要保存两个值的下标。（最小值、次小值）
*/
func minFallingPathSumII2(grid [][]int) int {
	n := len(grid)

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			var mdp int = math.MaxInt32

			for k := 0; k < n; k++ { // 枚举出了当前列之外的所有列
				if k == j {
					continue
				}
				if grid[i+1][k] < mdp {
					mdp = grid[i+1][k]
				}
			}
			grid[i][j] += mdp
		}
	}
	var ans int = math.MaxInt32
	for i := 0; i < n; i++ {
		if ans > grid[0][i] {
			ans = grid[0][i]
		}
	}
	return ans
}
