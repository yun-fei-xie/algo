package _0_dp

import "algo/leetcode/10_dp/kanpsack"

/*
https://leetcode.cn/problems/fibonacci-number/description/
*/
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

var fibMem = [31]int{0, 1}

func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if fibMem[n] != 0 {
		return kanpsack.mem[n]
	}
	fibMem[n] = fib2(n-1) + fib2(n-2)
	return fibMem[n]
}

/*
动态规划
*/
func fib3(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
