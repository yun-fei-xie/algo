package jianzhi_offer

/*
剑指 Offer 10- II. 青蛙跳台阶问题
https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/description/
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
示例 1：
*/
func numWays1(n int) int {
	var dfs func(x int) int
	dfs = func(x int) int {
		if x == 0 || x == 1 {
			return 1
		}
		return dfs(x-1) + dfs(x-2)
	}
	return dfs(n)
}

/*
动态规划解法
*/

func numWays2(n int) int {

	mod := 1000000007

	if n == 0 || n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod

	}

	return dp[n]
}
