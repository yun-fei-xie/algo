package sectionDP_test

import "math"

/*
2713. 矩阵中严格递增的单元格数
https://leetcode.cn/problems/maximum-strictly-increasing-cells-in-a-matrix/description/

给你一个下标从 1 开始、大小为 m x n 的整数矩阵 mat，你可以选择任一单元格作为 起始单元格 。
从起始单元格出发，你可以移动到 同一行或同一列 中的任何其他单元格，但前提是目标单元格的值 严格大于 当前单元格的值。
你可以多次重复这一过程，从一个单元格移动到另一个单元格，直到无法再进行任何移动。
请你找出从某个单元开始访问矩阵所能访问的 单元格的最大数量 。
返回一个表示可访问单元格最大数量的整数。

方法：
0.这个题和1340.maxJumps有点像，maxJumps在一维空间中进行跳跃，本题是在二维空间中进行跳跃。本质应该是一样的。
1.1340使用了一维的线性dp,这个题需要再两个方向上进行dp
2.base case mat[i][j]跳不动是因为这个值是第i行最大的，也是第j列最大的。为了能快速判断当前位置是否已经到达了base case 应该存储每一行每一列的最大值。
3.记忆化 558/564
*/
func maxIncreasingCells(mat [][]int) int {
	row, col := len(mat), len(mat[0])
	mem := make([][]int, row)
	for i := 0; i < row; i++ {
		mem[i] = make([]int, col)
		for j := 0; j < col; j++ {
			mem[i][j] = -1
		}
	}
	// 统计行、列的最大值
	rowMax := make([]int, row)
	colMax := make([]int, col)

	for i := 0; i < row; i++ {
		rowMax[i] = max(mat[i]...)
	}
	for j := 0; j < col; j++ {
		colM := math.MinInt
		for i := 0; i < row; i++ {
			if mat[i][j] > colM {
				colM = mat[i][j]
			}
		}
		colMax[j] = colM
	}

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
		if mat[i][j] == rowMax[i] && mat[i][j] == colMax[j] {
			return 1
		}
		//枚举行
		for k := 0; k < col; k++ {
			if k == i {
				continue
			}
			if mat[i][k] > mat[i][j] {
				ret = max(ret, dfs(i, k)+1)
			}
		}
		// 枚举列
		for k := 0; k < row; k++ {
			if k == i {
				continue
			}
			if mat[k][j] > mat[i][j] {
				ret = max(ret, dfs(k, j)+1)
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

/*
1:1翻译成动态规划
但是这个问题，我不知道状态的转移路径
*/
