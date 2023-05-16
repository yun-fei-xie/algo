package kanpsack

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/coin-change/?favorite=2cktkvj

零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

方法：
1.贪心算法 先兑换面值大的 (贪心不行) 因为兑换完大面额的，可会会出现无解
2.暴力枚举，当然，每一个面值的枚举的数量都有一个上限。这个上限就是当前余额amount 除以 当前硬币的面值，就得到一个最大的枚举数量
暴力枚举会超时。 是否可以进行优化->记忆化搜索
3.完全背包问题。（和0-1背包问题不同是，在0-1背包中，每个物品只有选择与不选择，而完全背包中会有不选择、选择1次或者多次）



*/

func coinChange2(coins []int, amount int) int {
	var res = math.MaxInt64
	var dfs func(coinIndex int, remainValue int, coinAmount int)
	dfs = func(coinIndex int, remainValue int, coinAmount int) {
		if coinIndex >= len(coins) || remainValue == 0 {
			if remainValue == 0 && coinAmount < res { // 找到了一个解
				res = coinAmount
			}
			return
		}

		// 一枚一枚地尝试
		// i -> [0...remainValue/coinValue]  i表示本次使用的硬币的枚数
		for i := 0; i <= remainValue/coins[coinIndex]; i++ {
			dfs(coinIndex+1, remainValue-i*coins[coinIndex], coinAmount+i)
		}

	}

	dfs(0, amount, 0)
	if res != math.MaxInt64 {
		return res
	}
	return -1
}

/*
递归
coinChange3这种写法已经可以使用记忆化
这里就不写了，直接在coinChange4中将代码翻译成dp
*/
func coinChange3(coins []int, amount int) int {
	// i表示考虑coins[0...i]这个区间，恰好兑换金额为c的面额需要的最少硬币数量
	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 0
			} else { // 无解，题目求最小值，这里返回一个最大值，上层在调用时，取min的时候，这个值就不会取到。
				return math.MaxInt32
			}
		}

		// 当前硬币coins[i]可以取0次到 c/coins[i]次。例如，当前剩余面额为10的时候，coins[i]=3,那么coins[i]最多使用3枚。
		var m = math.MaxInt32
		for j := 0; j <= c/coins[i]; j++ {
			m = min(dfs(i-1, c-j*coins[i])+j, m)
		}
		return m
	}

	return dfs(len(coins)-1, amount)
}

/*
动态规划
将coinChange3翻译成动态规划写法
*/
func coinChange4(coins []int, amount int) int {

	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		if i == 0 {
			dp[0][0] = 0
			for j := 1; j < amount+1; j++ {
				dp[0][j] = math.MaxInt32
			}
		}
	}

	// 递推
	for i := 1; i < len(dp); i++ {
		for c := 0; c < amount+1; c++ {
			var m = math.MaxInt32
			for j := 0; j <= c/coins[i-1]; j++ {
				m = min(m, dp[i-1][c-j*coins[i-1]]+j)
			}
			dp[i][c] = m
		}
	}
	if dp[len(dp)-1][amount] == math.MaxInt32 {
		return -1
	}
	return dp[len(dp)-1][amount]
}

func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}

/*
错误解法: 这个题不能用贪心算法
*/
func coinChange1(coins []int, amount int) int {
	res := 0

	sort.Slice(coins, func(i, j int) bool {
		if coins[i] > coins[j] {
			return true
		}
		return false
	})

	for i := 0; i < len(coins); {
		value := coins[i]
		if value > amount { // 当前这个面值大了，需要换个小点的进行兑换
			i++
			continue
		} else {
			cnt := amount / value
			res += cnt
			amount -= cnt * value
			if amount == 0 {
				break
			}
			i++
		}
	}

	if amount != 0 { // 兑换不完
		return -1
	} else {
		return res
	}

}

func TestCoinChange(t *testing.T) {
	//fmt.Println(coinChange([]int{1, 2, 5}, 11))
	//fmt.Println(coinChange([]int{2}, 3))
	//fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange2([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange3([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange4([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange2([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864))
	fmt.Println(coinChange3([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864))
	fmt.Println(coinChange4([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864))
}
