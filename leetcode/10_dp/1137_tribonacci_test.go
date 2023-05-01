package _0_dp

/*
https://leetcode.cn/problems/n-th-tribonacci-number/?envType=study-plan-v2&id=dynamic-programming
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
