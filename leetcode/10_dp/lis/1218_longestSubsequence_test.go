package lis

import (
	"fmt"
	"testing"
)

func longestSubsequence1(arr []int, difference int) int {
	var ans int = 1
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == 0 {
			return 1
		}
		m := 1
		for j := i - 1; j >= 0; j-- {
			// 必须先求解一下子问题
			cnt := dfs(j)
			if arr[i]-arr[j] == difference {
				m = max(m, cnt+1)
			}
		}
		ans = max(ans, m)
		return m
	}
	dfs(len(arr) - 1)
	return ans
}

/*
动态规划：翻译成递推  超时-😅 还能继续优化
*/
func longestSubsequence2(arr []int, difference int) int {
	var ans int
	dp := make([]int, len(arr))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(dp); i++ {
		var m int = 1
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] == difference {
				m = max(m, dp[j]+1)
			}
		}
		dp[i] = m
		ans = max(ans, dp[i])
	}
	return ans
}

/*
借用两个数之和的思想继续优化
如果我们知道了以arr[i]为结尾的最长等差数列的长度为l
用一个map[int]int保存，这个值。

我们遍历到arr[i]的时候，需要通过遍历arr[j]->[0<=j<=i-1]这么多数字，看看哪个数字等于arr[i]-diff =arr[j]
通过使用hash表，可以降低查找的时间复杂度。
*/
func longestSubsequence3(arr []int, difference int) int {
	var ans int
	dp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		dp[arr[i]] = dp[arr[i]-difference] + 1
		if dp[arr[i]] > ans {
			ans = dp[arr[i]]
		}
	}
	return ans
}

func TestLongestSubSequence(t *testing.T) {
	fmt.Println(longestSubsequence1([]int{3, 4, -3, -2, -4}, -5))
	fmt.Println(longestSubsequence2([]int{3, 4, -3, -2, -4}, -5))
	fmt.Println(longestSubsequence1([]int{-13, 26, -4, -1, -2, -28, 21, 8, -26, 5, -20, -19, -1, 1, 19, 29, -9, -8, 5, 9, -29, 18, 14, -24, 24, 26, -6, -26, -19}, 10))
	fmt.Println(longestSubsequence2([]int{-13, 26, -4, -1, -2, -28, 21, 8, -26, 5, -20, -19, -1, 1, 19, 29, -9, -8, 5, 9, -29, 18, 14, -24, 24, 26, -6, -26, -19}, 10))
}
