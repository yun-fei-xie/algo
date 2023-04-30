package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/minimum-path-sum/

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
*/

/*
递归暴力搜索
每次有两种走法，起点和终点固定（走到终点递归终止）
*/
func minPathSum(grid [][]int) int {

	row := len(grid)
	col := len(grid[0])
	var ans = math.MaxInt64
	var dfs func(grid [][]int, x int, y int, sum int)
	dfs = func(grid [][]int, x int, y int, sum int) {
		// 越界
		if x >= row || y >= col {
			return
		}
		// 达到目标地点
		if x == row-1 && y == col-1 {
			sum += grid[x][y]
			if sum < ans {
				ans = sum
			}
			return
		}

		dfs(grid, x+1, y, sum+grid[x][y])
		dfs(grid, x, y+1, sum+grid[x][y])
	}
	dfs(grid, 0, 0, 0)
	return ans
}

/*
我只要知道当前点向右走一步到终点的最短距离 和
倒着往上推 需要控制好索引的顺序。这个顺序就是自底向上，每一行，从右向左（最下一行和最右列 初始直接赋值，这样可以避免考虑边界条件）
这个题不能扩展数组，扩充数组会在min函数的时候会产生逻辑错误
*/

func miniPathSum2(grid [][]int) int { // m * n

	var min func(x, y int) int
	min = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	m := len(grid)
	n := len(grid[0])

	//m * n
	dp := make([][]int, 0)
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))

	}

	cum := 0
	for j := len(dp[0]) - 1; j >= 0; j-- { // 赋值最后一行
		cum = cum + grid[len(dp)-1][j]
		dp[len(dp)-1][j] = cum
	}

	cum = 0
	for j := len(dp) - 1; j >= 0; j-- { // 赋值最右列
		cum += grid[j][len(dp[0])-1]
		dp[j][len(dp[0])-1] = cum
	}

	// 递推
	for i := len(dp) - 2; i >= 0; i-- {
		for j := len(dp[i]) - 2; j >= 0; j-- {
			dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
		}
	}
	return dp[0][0]
}

func TestPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	res := miniPathSum2(grid)
	fmt.Println(res)

	grid2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	res2 := miniPathSum2(grid2)
	fmt.Println(res2)

}
