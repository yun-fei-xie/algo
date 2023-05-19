package kanpsack

import (
	"fmt"
	"sort"
	"testing"
)

/*
377. 组合总和 Ⅳ
https://leetcode.cn/problems/combination-sum-iv/description/
扩展一下：返回所有的可能结果
这一版代码求出的是真的组合，题目描述的更像是排列。
假设target=6，先循环物品，肯定不会有{1,2,3}和{2,3,1}同时出现。
因为，每次都是从一个一个点朝着一个方向进行遍历，无法做到在固定的背包容量下遍历所有物品。

如果先循环背包容量，再循环物品。就可以在内层循环中，遍历所有的物品。
*/
func combinationSum4(nums []int, target int) int {
	// 将物品的数量从小到大进行排序
	sort.Ints(nums)
	// [0...i]这个区间符合要求的组合数量
	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}

		// c的容量可以装下当前物品
		var ret int
		for j := 0; j <= c/nums[i]; j++ {
			ret += dfs(i-1, c-nums[i]*j)
		}
		return ret
	}

	return dfs(len(nums)-1, target)
}

/*
正确的解法
*/
func combinationSum2(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			// 如果一个都装不进去 说明凑不出这个target
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

func TestCombinationSum4(t *testing.T) {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	//fmt.Println(combinationSum4([]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}, 10))
}
