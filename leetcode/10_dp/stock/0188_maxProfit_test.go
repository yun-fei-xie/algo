package stock

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/


给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格，和一个整型 k 。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2：

输入：k = 2, prices = [3,2,6,5,0,3]
输出：7
解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
*/

/*
递归解法
参考 https://www.bilibili.com/video/BV1ho4y1W7QK/?spm_id_from=333.999.0.0&vd_source=4bc6e42744318f1802754628d7048e6d
*/
func maxProfit(k int, prices []int) int {
	// i -> prices[0...i] 这个区间考虑第i个价格  j->最多还能做多少笔交易  hold->当前是否持有股票
	var dfs func(i int, j int, hold bool) int
	dfs = func(i int, j int, hold bool) int {
		if j < 0 { // j表示当前最多能做多少笔交易
			return math.MinInt32
		}
		if i < 0 {
			if hold {
				return math.MinInt32
			}
			return 0
		}
		if hold {
			return max(dfs(i-1, j, true), dfs(i-1, j-1, false)-prices[i])
			// 当前能做j笔交易，那么i-1没有持有股票->买入 就只能做j-1次交易
		}
		return max(dfs(i-1, j, false), dfs(i-1, j, true)+prices[i])

	}
	return dfs(len(prices)-1, k, false)
}

/*
记忆化
i->[0...len(prices)-1]
j->[0...k]
dfs有三个状态参数，i,j,hold
*/

func maxProfit2(k int, prices []int) int {
	mem := make([][][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		mem[i] = make([][2]int, k+1) // j 的取值范围[0...k]
		for j := 0; j <= k; j++ {
			mem[i][j] = [...]int{-1, -1}
		}
	}

	// i -> prices[0...i] 这个区间考虑第i个价格  j->最多还能做多少笔交易  hold->当前是否持有股票
	var dfs func(i int, j int, hold int) (res int)
	dfs = func(i int, j int, hold int) (res int) {
		if j < 0 { // j表示当前最多能做多少笔交易
			return math.MinInt32
		}
		if i < 0 {
			if hold == 1 {
				return math.MinInt32
			}
			return 0
		}
		// 记忆化 这个地方的写法以后多用用
		if m := mem[i][j][hold]; m != -1 {
			return m
		}
		defer func() {
			mem[i][j][hold] = res
		}()

		if hold == 1 {
			return max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-prices[i])
			// 当前能做j笔交易，那么i-1没有持有股票->买入 就只能做j-1次交易
		}
		return max(dfs(i-1, j, 0), dfs(i-1, j, 1)+prices[i])

	}
	return dfs(len(prices)-1, k, 0)
}

/*
动态规划：递推
*/

func maxProfit3(k int, prices []int) int {

	// i->[0...len(prices)-1]  j->[0...k] , stock->[0...1]
	dp := make([][][2]int, len(prices)+1)
	for i := 0; i < len(prices)+1; i++ {
		dp[i] = make([][2]int, k+2)
		for j := 0; j <= k+1; j++ {
			dp[i][j] = [2]int{math.MinInt32, math.MinInt32}
		}
	}
	// i<0并且不持有股票
	for j := 1; j <= k+1; j++ {
		dp[0][j][0] = 0
	}

	for i := 1; i <= len(prices); i++ {
		for j := 1; j <= k+1; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1]) // 卖出不消耗交易次数dp[i-1][j][1]+prices[i-1] 这里写j
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}

	return dp[len(dp)-1][k+1][0]
}

func TestMaxProfitIV(t *testing.T) {
	//fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
	fmt.Println(maxProfit2(2, []int{3, 2, 6, 5, 0, 3}))
	fmt.Println(maxProfit3(2, []int{3, 2, 6, 5, 0, 3}))
}
