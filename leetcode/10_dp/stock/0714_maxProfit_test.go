package stock

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
返回获得利润的最大值。
注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
示例 1：
输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
输出：8
解释：能够达到的最大利润:
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8

这个题和122基本一致 就多了一个手续费 买入的时候把手续费扣掉就行

*/

/*
递归解法
*/
func maxProfitFee(prices []int, fee int) int {

	var dfs func(i int, hold int) int
	dfs = func(i int, hold int) int {
		if i < 0 {
			if hold == 1 {
				return math.MinInt32
			}
			return 0
		}
		if hold == 1 {
			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i]+fee)
		}
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
	}
	return dfs(len(prices)-1, 0)
}

/*
动态规划
*/
func maxProfitFeeDp(prices []int, fee int) int {

	// 本题只有2个状态
	dp := make([][2]int, len(prices)+1)
	dp[0][1] = math.MinInt32
	for i := 1; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1]-fee)
	}
	return dp[len(dp)-1][0]
}

func TestMaxProfitFee(t *testing.T) {
	fmt.Println(maxProfitFeeDp([]int{1, 3, 2, 8, 4, 9}, 2))
}
