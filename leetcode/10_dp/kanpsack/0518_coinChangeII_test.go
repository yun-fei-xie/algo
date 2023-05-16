package kanpsack

import (
	"fmt"
	"testing"
)

/*
518. 零钱兑换 II
https://leetcode.cn/problems/coin-change-ii/

给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。
题目数据保证结果符合 32 位带符号整数。

示例 1：
输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

方法：本质是一个完全背包问题。
求的是方案数。


*/
/*
递归写法
*/
func change1(amount int, coins []int) int {

	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}

		var cnt int
		for j := 0; c-j*coins[i] >= 0; j++ {
			cnt += dfs(i-1, c-j*coins[i])
		}
		return cnt
	}
	return dfs(len(coins)-1, amount)
}

/*
将chang1的递归翻译成动态规划
*/
func change2(amount int, coins []int) int {

	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1

	for i := 1; i < len(dp); i++ {
		for c := 0; c < len(dp[0]); c++ {
			var cnt int
			for j := 0; c-j*coins[i-1] >= 0; j++ {
				cnt += dp[i-1][c-j*coins[i-1]]
			}
			dp[i][c] = cnt
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func TestChangeCoinII(t *testing.T) {
	fmt.Println(change1(5, []int{1, 2, 5}))
	fmt.Println(change2(5, []int{1, 2, 5}))
	fmt.Println(change1(3, []int{2}))
	fmt.Println(change2(3, []int{2}))
}
