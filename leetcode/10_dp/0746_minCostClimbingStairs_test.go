package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/min-cost-climbing-stairs/

这道题需要注意的一个点是，需要跳出数组外才算结束！
*/
func minCostClimbingStairs(cost []int) int {
	memo := make([]int, len(cost)+1)
	memo[0] = 0
	memo[1] = 0 // 前两级台阶不费力气

	for i := 2; i < len(memo); i++ {
		memo[i] = int(math.Min(float64(memo[i-1]+cost[i-1]), float64(memo[i-2]+cost[i-2])))
		// 后一级台阶有两种方式可以跳过来，选择消耗少的

	}
	return memo[len(memo)-1]
}

func TestMinCostClimbingStairs(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))

}
