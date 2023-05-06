package _0_dp

/*
https://leetcode.cn/problems/n-th-tribonacci-number/?envType=study-plan-v2&id=dynamic-programming

泰波那契序列 Tn 定义如下：
T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
*/
func tribonacci(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	if n <= 2 {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
