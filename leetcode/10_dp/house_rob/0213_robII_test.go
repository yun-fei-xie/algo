package house_rob

import (
	"fmt"
	"testing"
)

/*
213. 打家劫舍 II
https://leetcode.cn/problems/house-robber-ii/

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

和198不同之处在于，数组首位相连，如何判断是否已经达到了区间的末尾？(走到什么地步不能走了)

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

方法：两次动态规划
这道题和第918号问题，环形连续子数组最大和有点像。
它多了一个条件。相邻的元素不能同时出现。
如何才能保证第一间房屋和最后一间房屋不同时偷窃呢？
如果偷窃了第一间房屋，则不能偷窃最后一间房屋，因此偷窃房屋的范围是第一间房屋到最后第二间房屋；
也就是考虑偷窃[0...n-2]。
如果偷窃了最后一间房屋，则不能偷窃第一间房屋，因此偷窃房屋的范围是第二间房屋到最后一间房屋。
也就是考虑偷窃[1...n-1]。
最后的结果取两者的最大值。
于是这个问题被规约为做2次打家劫舍问题。
*/

func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m1 := houseRob(nums, 1, len(nums)-1)
	m2 := houseRob(nums, 0, len(nums)-2)
	return max(m1, m2)
}
func houseRob(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	dp := make([]int, len(nums))
	dp[left] = nums[left]
	dp[left+1] = max(nums[left], nums[left+1])
	for i := left + 2; i <= right; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[right]
}

func TestRob2(t *testing.T) {
	fmt.Println(robII([]int{1, 2, 3, 1}))
	fmt.Println(robII([]int{2, 3, 2}))
	fmt.Println(robII([]int{1, 2, 3}))
}
