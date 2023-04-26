package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/min-cost-climbing-stairs/

这道题需要注意的一个点是，需要跳出数组外才算结束！
新题目描述中明确说了 “你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。” 也就是说 从 到达 第 0 个台阶是不花费的，但从 第0 个台阶 往上跳的话，需要花费 cost[0]。
所以初始化 dp[0] = 0，dp[1] = 0;
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

/*
深度优先搜索
*/
func minCostClimbingStairs2(cost []int) int {
	var minCost = math.MaxInt64
	var traceback func(startIndex int, totalCost int)
	traceback = func(startIndex int, totalCost int) {
		if startIndex >= len(cost) { // 超过了台阶的数量->爬完
			if totalCost < minCost {
				minCost = totalCost
			}
			return
		}
		// 来到当前这一级
		traceback(startIndex+1, totalCost+cost[startIndex]) //向上爬一级
		traceback(startIndex+2, totalCost+cost[startIndex]) // 向上爬两级
	}
	traceback(0, 0) //从第0级开始爬
	traceback(1, 0) //从第1级开始爬
	return minCost
}

// 如何使用记忆化搜索？

func TestMinCostClimbingStairs(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	cost2 := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
	fmt.Println(minCostClimbingStairs2(cost))

	fmt.Println(minCostClimbingStairs(cost2))
	fmt.Println(minCostClimbingStairs2(cost2))

}
