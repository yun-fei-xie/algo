package stock

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/?favorite=2cktkvj
题目中的条件：在这个周期中，只能买入和卖出一次
*/

func maxProfitI1(prices []int) int {
	// 只能有一次买卖吗？  只能有一次买卖
	minPrice := math.MaxInt64
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxProfit = max(maxProfit, prices[i]-minPrice)
	}
	return maxProfit
}

// 暴力解法 双重循环
func maxProfitI2(prices []int) int {
	var maxProfit int
	var length = len(prices)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[i]
			}
		}
	}
	return maxProfit
}

// 贪心算法 靠，这和上面写的是一样的
func maxProfitI3(prices []int) int {
	// 只能有一次买卖吗？  只能有一次买卖
	minPrice := math.MaxInt64
	ans := 0
	dp := make([]int, len(prices))

	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		dp[i] = prices[i] - minPrice
		ans = max(ans, dp[i])
	}
	return ans
}

func maxProfitI4(prices []int) int {
	dp := make([][2]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = [2]int{}
	}
	// dp[i][0] -> 第i天结束的时候，持有0份股票 最大利润
	// dp[i][1] -> 第i天结束的时候，持有1份股票 最大利润

	// 最应该注意的地方时，如果第i天持有，那么只有两种情况。
	// 1.[0...i-1]某一天买入，但是一直没卖掉。 2.今天买入
	dp[0][0] = 0
	dp[0][1] = -prices[0] //买入
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		//dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[len(prices)-1][0]
}

//func maxProfitI5(prices []int) int {
//	dp := make([][][]int) // dp[i][j][k]
//}

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfitI1([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfitI2([]int{7, 1, 5, 3, 6, 4}))
	//prices [7,1,5,3,6,4]
	//min    [7,1,1,1,1,1]
	//dp     [0,0,4,2,5,3]
}
