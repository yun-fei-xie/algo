package house_rob

import (
	"fmt"
	"math"
	"testing"
)

/*
198. 打家劫舍
题目链接:https://leetcode.cn/problems/house-robber/

你是一个专业的小偷，计划偷窃沿街的房屋。
每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

方法：递归+记忆化搜索
对于数组类的问题，有两种思考方式，一种是从左到右进行思考。
第i个元素的状态依赖前[0...i-1]的状态。
最简单的案例就是斐波那契数列的求解。f(i)= f(i-1)+ f(i-2)
具体代码参考rob2。

方法：动态规划
将递归1：1翻译成递推。具体代码和解释参考rob3
*/
func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	var dfs func(nums []int, startIndex int) int
	dfs = func(nums []int, startIndex int) int {
		if startIndex >= len(nums) { // 没房子可以打劫 直接返回
			return 0
		}
		if memo[startIndex] != -1 {
			return memo[startIndex]
		}

		// [startIndex , n-1] 扫荡这个区间 目的是求解打劫这个区间能够获得的最大收益
		max := math.MinInt64
		for i := startIndex; i < len(nums); i++ {
			m := nums[i] + dfs(nums, i+2)
			if m > max {
				max = m
			}
		}
		memo[startIndex] = max

		return memo[startIndex]
	}

	return dfs(nums, 0)
}

/*
从左向右进行递归类似斐波那契数列求解
dfs(i int)int ,i表示当前考虑打劫第i间房子。
它返回的是在[0...i]这么多间房子中挑选一些房子进行打劫，可以获得的最大收益。
需要注意的是当前的房子可以打劫，也可以不打劫。
它的最大值取决于
*/

func rob2(nums []int) int {
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		return max(dfs(i-2)+nums[i], dfs(i-1))
	}
	return dfs(len(nums) - 1)
}

/*
将rob2这份递归代码1：1翻译成递推。
递归函数中只有一个参数，那么dp数组应该是1维的。
递归的边界条件就应该是dp数组的临界条件。
在递归中用到了dfs(i-2) ，那么就需要考虑dp[0]和dp[1]如何初始化。
 1. 给dp数组多预留两个位置。在计算的时候，使用下标平移的方式。
    多出来位置的初始值赋值为递归函数中越界的值。这个题目是 if i<0 -> return 0。
 2. 直接初始化dp[0],dp[1]。在循环计算的时候，下标从2开始。
*/
func rob3(nums []int) int {
	dp := make([]int, len(nums)+2)

	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-2])
	}
	return dp[len(dp)-1]
}

/*
dp解法：自底向上进行推导，先从最后开始推
*/
func robdp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	memo := make([]int, len(nums))
	memo[len(nums)-1] = nums[len(nums)-1]

	for i := len(nums) - 1; i >= 0; i-- { // 每一个节点
		// [i...n-1]
		max := math.MinInt64
		for j := i; j < len(nums); j++ {
			var m int
			if j+2 >= len(nums) { // 如果越界-> 对应递归到底 到达一个空的位置 此时直接返回0
				m = nums[j] + 0
			} else {
				m = nums[j] + memo[j+2]
			}
			if m > max {
				max = m
			}
		}
		memo[i] = max
	}
	return memo[0]
}

/*
最简洁的dp
dp[i]表示在[0...i]这么多间房子进行偷窃可以获得的最大金额。
对于第i间房子来说，如果这间房子要偷的话，第i-1间房子就不能偷。
于是问题转换成了nums[i]+dp[i-2]。
如果第i间房子不偷，那么问题转化成了[0...i-1]间房子的最大值。
对于dp[i]来说，它应该等于两者的较大值 dp[i] = max{nums[i]+dp[i-2] , dp[i-1]}
*/
func robDp2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) < 2 {
		return nums[0]
	}
	dp[1] = max(nums[0], nums[1])

	for j := 2; j < len(nums); j++ {
		dp[j] = max(nums[j]+dp[j-2], dp[j-1])
	}
	return dp[len(nums)-1]
}

func TestRob(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	// dp        [2, 7, 11,11,12 ]
	fmt.Println(rob(nums))
	fmt.Println(robdp(nums))
	fmt.Println(robDp2(nums))
}
