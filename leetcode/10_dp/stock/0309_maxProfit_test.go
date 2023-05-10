package stock

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/

给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格。
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入: prices = [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
示例 2:
输入: prices = [1]
输出: 0
*/
func maxProfit1(prices []int) int {
	// 在prices[0...index]这个区间进行交易，可以获得的最大收益。 含有冷冻期
	var dfs func(index int, hasStock bool) int
	dfs = func(index int, hasStock bool) int {
		if index < 0 {
			if hasStock {
				return math.MinInt32
			}
			return 0
		}

		if hasStock {
			return max(dfs(index-1, hasStock), dfs(index-2, false)-prices[index])
		}
		return max(dfs(index-1, false), dfs(index-1, true)+prices[index])
	}
	return dfs(len(prices)-1, false)
}

/*
如何保证冷冻 打家劫舍问题 不相邻
*/
func maxProfitDp(prices []int) int {

	dp := make([][2]int, len(prices)+2) // 预留空位给边界条件
	for i := 0; i < len(dp); i++ {
		dp[i] = [2]int{}
	}
	dp[1][1], dp[0][1] = math.MinInt32, math.MinInt32

	for i := 2; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-2])
		//当前不持有股票
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i-2])
	}
	return dp[len(dp)-1][0]
}

func TestMaxProfitCoolDown(t *testing.T) {
	fmt.Println(maxProfit1([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfitDp([]int{1, 2, 3, 0, 2}))
}
