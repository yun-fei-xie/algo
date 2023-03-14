package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
01背包问题的代码
*/

func knapsack01(weight []int, values []int, c int) int {

	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	dp := make([][]int, 0)
	for i := 0; i < len(weight); i++ { // 行数->物品的数量
		dp = append(dp, make([]int, c+1)) // dp数组的列 长度为[0...c]背包重量这样的长度
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 0 //背包容量为0
	}

	for i := 1; i <= c; i++ {
		if weight[0] > c {
			dp[0][i] = 0
		} else {
			dp[0][i] = values[0]
		}
	}

	// 目标是dp数组右下角那个数字
	for i := 1; i < len(weight); i++ {
		// 每一件物品i 放还是不放
		// 如何遍历j-> j从下到大因为现在下面的下标中，dp[i-1][j-weight[i]] 需要用到
		for j := 1; j <= c; j++ { // 遍历背包容量(j是下标 也是背包的容量)
			if weight[i] > j { // 放不下
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+values[i])
			}
		}
	}

	return dp[len(weight)-1][c]
}

/*
暴力解法：枚举
*/
func knapsackEnum(weight []int, values []int, c int) int {

	var dfs func(startIndex int, cap int, value int)
	maxValue := math.MinInt64

	//startIndex ：当前待考察的物品下标
	dfs = func(startIndex int, cap int, value int) {
		if cap <= 0 || startIndex >= len(weight) {

			return
		}

		if cap >= weight[startIndex] { // 还有放的位置 就放一个进去试试
			if value+values[startIndex] > maxValue { // 放完之后尝试更新当前背包中的最大值
				maxValue = value + values[startIndex]
			}
			dfs(startIndex+1, cap-weight[startIndex], value+values[startIndex])
		}

		dfs(startIndex+1, cap, value) // 不放置当前元素肯定不需要放置

	}
	dfs(0, c, 0)

	return maxValue
}

func TestKnapsack(t *testing.T) {
	weights := []int{1, 3, 4}
	values := []int{15, 20, 30}
	c := 5
	res := knapsack01(weights, values, c)
	fmt.Println(res)

	res2 := knapsackEnum(weights, values, c)
	fmt.Println(res2)
}
