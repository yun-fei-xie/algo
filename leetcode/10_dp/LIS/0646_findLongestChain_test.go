package LIS

import (
	"fmt"
	"sort"
	"testing"
)

/*
最长递增子序列问题
*/
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] < pairs[j][0] {
			return true
		}
		return pairs[i][1] < pairs[j][1]
	})

	//dp[i]	 -> 以pairs[i]为结尾的最长自增子序列的长度
	var dp = make([]int, len(pairs))
	var ans = 1
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return dp[len(pairs)-1]
}

func TestFindLongestChain(t *testing.T) {
	fmt.Println(findLongestChain([][]int{{1, 2}, {7, 8}, {4, 5}}))
}
