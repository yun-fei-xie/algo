package kanpsack

/*
474. 一和零
https://leetcode.cn/problems/ones-and-zeroes/description/

给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

方法：转换成0-1背包问题。
虽然有两种物品，但是放在一起考虑就没有问题
*/
func findMaxForm(strs []string, m int, n int) int {

	var dfs func(i int, c1 int, c2 int) int
	//[i...len(s)-1]最大子集
	dfs = func(i int, c1 int, c2 int) int {
		if i >= len(strs) {
			return 0
		}
		// 考虑i 可以装 或者不装
		zero := getZero(strs[i])
		one := len(strs[i]) - zero
		var maxSub = 0
		// 如果能装
		if zero <= c1 && one <= c2 {
			maxSub = max(maxSub, dfs(i+1, c1-zero, c2-one)+1)
		}
		// 如果装不下
		maxSub = max(maxSub, dfs(i+1, c1, c2))
		// 返回二者最大值
		return maxSub
	}
	return dfs(0, m, n)
}
func max(num ...int) int {
	m := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > m {
			m = num[i]
		}
	}
	return m
}

func getZero(s string) (ans int) {
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			ans++
		}
	}
	return ans
}

/*
1:1翻译成动态规划
*/
func findMaxForm2(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := len(strs) - 1; i >= 0; i-- {
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				// i->当前考虑的字符串下标  j->当前0的剩余容量  k->当前1的剩余容量
				var maxSub = 0
				// 考虑i 可以装 或者不装
				zero := getZero(strs[i])
				one := len(strs[i]) - zero
				if j >= zero && k >= one {
					maxSub = max(maxSub, dp[i+1][j-zero][k-one]+1)
				}
				maxSub = max(maxSub, dp[i+1][j][k])
				dp[i][j][k] = maxSub
			}
		}
	}
	return dp[0][m][n]
}
