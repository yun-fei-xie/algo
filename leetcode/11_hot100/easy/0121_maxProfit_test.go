package easy

import "math"

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/?favorite=2cktkvj
题目中的条件：在这个周期中，只能买入和卖出一次
*/

func maxProfit(prices []int) int {
	// 只能有一次买卖吗？  只能有一次买卖
	minPrice := math.MaxInt64
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else { // 不是最低价格
			if prices[i]-minPrice > maxProfit {
				maxProfit = prices[i] - minPrice
			}
		}
	}
	return maxProfit
}
