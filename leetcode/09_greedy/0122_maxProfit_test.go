package _9_greedy

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

假如第0天买入，第3天卖出，那么利润为：prices[3] - prices[0]。
相当于(prices[3] - prices[2]) + (prices[2] - prices[1]) + (prices[1] - prices[0])。
此时就是把利润分解为每天为单位的维度，而不是从0天到第3天整体去考虑！
*/
func maxProfit(prices []int) int {
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

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Println(profit)
}
