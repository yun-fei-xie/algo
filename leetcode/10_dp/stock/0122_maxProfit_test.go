package stock

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
返回 你能获得的 最大 利润 。
示例 1：

输入：prices = [7,1,5,3,6,4]
输出：7
解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
     总利润为 4 + 3 = 7 。
示例 2：

输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
     总利润为 4 。
示例 3：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0 。


假如第0天买入，第3天卖出，那么利润为：prices[3] - prices[0]。
相当于(prices[3] - prices[2]) + (prices[2] - prices[1]) + (prices[1] - prices[0])。
此时就是把利润分解为每天为单位的维度，而不是从0天到第3天整体去考虑！
*/

// 解法1 暴力枚举
func maxProfitII1(prices []int) int {
	var maxProfit int
	var dfs func(index int, hasStock bool, profit int)
	dfs = func(index int, hasStock bool, profit int) {
		if index == len(prices) { // 递归到底
			maxProfit = max(maxProfit, profit)
			return
		}
		// 什么都不做
		dfs(index+1, hasStock, profit)

		if hasStock {
			// 手上有股票->卖出 卖出挣钱所以是+
			dfs(index+1, !hasStock, profit+prices[index])
		} else {
			// 手上没有股票->买入 买入花钱所以是-
			dfs(index+1, !hasStock, profit-prices[index])
		}
	}
	dfs(0, false, 0)
	return maxProfit
}

func maxProfitII2(prices []int) int {
	profit := make([]int, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		profit[i-1] = prices[i] - prices[i-1]
	}
	count := 0
	for i := 0; i < len(profit); i++ {
		if profit[i] > 0 {
			count += profit[i]
		}
	}
	return count
}

// 解法3 动态规划
func maxProfitII3(prices []int) int {
	dp := make([][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = [2]int{}
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(dp)-1][0]
}

/*
倒着递推，写法参考B站 灵茶山艾府
https://www.bilibili.com/video/BV1ho4y1W7QK/?spm_id_from=333.999.0.0&vd_source=4bc6e42744318f1802754628d7048e6d
*/
func maxProfitII4(prices []int) int {

	var dfs func(statIndex int, hasStock bool) int // [0...startIndex] 这个区间买卖股票可以持有的最大收益
	dfs = func(statIndex int, hasStock bool) int {
		if statIndex < 0 {
			if hasStock {
				return math.MinInt32
			} else {
				return 0
			}
		}

		if hasStock { // 当前持有股票
			return max(dfs(statIndex-1, true), dfs(statIndex-1, false)-prices[statIndex])
		} else { //当前不持有股票
			return max(dfs(statIndex-1, false), dfs(statIndex-1, true)+prices[statIndex])
		}

	}
	return dfs(len(prices)-1, false)
}

// 把倒着递推改成正向递推
// 只需要两个状态 i 第i天 j 是否持有股票
// dp[i][j]表示第i天的时候能够获得的最大收益

func maxProfitII5(prices []int) int {
	dp := make([][2]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = [2]int{}
	}
	dp[0][0], dp[0][1] = 0, math.MinInt32
	for i := 1; i < len(dp); i++ { //dp[i]->prices[i-1]
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}
	// return max(dp[len(prices)[0]] , dp[len(prices)][1])
	return dp[len(prices)][0] // 最后一天不持有股票 必然比持有股票收益更大
}

func TestMaxProfit2(t *testing.T) {
	fmt.Println(maxProfitII2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfitII1([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfitII3([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfitII4([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfitII5([]int{7, 1, 5, 3, 6, 4}))
}
