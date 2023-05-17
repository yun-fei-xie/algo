package sectionDP

import (
	"fmt"
	"math"
	"testing"
)

/*
375. 猜数字大小 II
https://leetcode.cn/problems/guess-number-higher-or-lower-ii/description/

我们正在玩一个猜数游戏，游戏规则如下：

我从 1 到 n 之间选择一个数字。
你来猜我选了哪个数字。
如果你猜到正确的数字，就会 赢得游戏 。
如果你猜错了，那么我会告诉你，我选的数字比你的 更大或者更小 ，并且你需要继续猜数。
每当你猜了数字 x 并且猜错了的时候，你需要支付金额为 x 的现金。如果你花光了钱，就会 输掉游戏 。
给你一个特定的数字 n ，返回能够 确保你获胜 的最小现金数，不管我选择那个数字 。

方法：递归
假设我区间[i...j]中选择了一个数字x，如果选择的不对，那么需要支付现金x,
并且需要左右区间dfs(i,x-1)->[i...x-1]和dfs(x+1 ,j)->[x+1...j]
中继续进行选择。为区间[i...j]支付的最小金额是 x+max{dfs(i,x-1),dfs(x+1,j)}，这是因为我们需要覆盖住最坏情况。
综合所有的最坏情况，选择一个最小值。

尝试选择区间[i...j]中的每一个x,计算dfs(i,j) -> 就有dfs(i,j) = min(x+max{dfs(i,x-1),dfs(x+1,j)})
*/
func getMoneyAmount1(n int) int {

	// 在[i...j]这个范围里面擦测一个数字，使得费用最小
	// 猜测的数字可以是k, i<=k<=j
	// 当i==j的时候 肯定能猜对，费用为0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}
		// j>i 区间至少有2个数字
		var m = math.MaxInt
		for k := i; k <= j; k++ {
			m = min(m, max(dfs(i, k-1)+k, dfs(k+1, j)+k))

		}
		//fmt.Printf("i->%d  j->%d  dfs(i,j)->%d\n", i, j, m)
		return m
	}
	return dfs(1, n)
}

/*
1:1动态规划
需要注意区间的范围
*/

func getMoneyAmount2(n int) int {

	dp := make([][]int, n+10)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+10) // 为了防止数组越界，两边留下富裕空间
	}

	// i 倒序遍历 j正序遍历
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			// 区间范围[i...j]保证区间长度最小是2
			var m = math.MaxInt
			for k := i; k <= j; k++ {
				m = min(m, max(dp[i][k-1], dp[k+1][j])+k)
			}
			dp[i][j] = m
		}
	}
	return dp[1][n]
}

func TestGetMoneyAmount(t *testing.T) {
	fmt.Println(getMoneyAmount1(10))
	fmt.Println(getMoneyAmount2(10))
}
