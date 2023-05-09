package stock

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。

	随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。

示例 2：
输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。

	注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
	因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

思路：这个题和188一样，只是将188题的k改为了2.
*/
func maxProfitIII(prices []int) int {

	var dfs func(i int, j int, hold int) int
	dfs = func(i int, j int, hold int) int {
		if j < 0 {
			return math.MinInt32
		}
		if i < 0 {
			if hold == 1 {
				return math.MinInt32
			}
			return 0
		}

		if hold == 1 {
			return max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-prices[i])
		}
		return max(dfs(i-1, j, 0), dfs(i-1, j, 1)+prices[i])
	}
	return dfs(len(prices)-1, 2, 0)
}

func maxProfitIIIDp(prices []int) int {

	dp := make([][][2]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, 4) // len([0,1,2])+1 -> 4
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = [2]int{math.MinInt32 / 2, math.MinInt32 / 2}
		}
	}

	for j := 1; j < 4; j++ {
		dp[0][j][0] = 0
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < 4; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}
	return dp[len(dp)-1][len(dp[0])-1][0]
}

func TestMaxProfitIII(t *testing.T) {
	fmt.Println(maxProfitIII([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfitIII([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfitIIIDp([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
