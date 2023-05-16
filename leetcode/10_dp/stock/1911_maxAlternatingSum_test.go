package stock

import (
	"fmt"
	"math"
	"testing"
)

/*
1911. 最大子序列交替和
https://leetcode.cn/problems/maximum-alternating-subsequence-sum/

一个下标从 0 开始的数组的 交替和 定义为 偶数 下标处元素之 和 减去 奇数 下标处元素之 和 。
比方说，数组 [4,2,5,3] 的交替和为 (4 + 5) - (2 + 3) = 4 。
给你一个数组 nums ，请你返回 nums 中任意子序列的 最大交替和 （子序列的下标 重新 从 0 开始编号）。
一个数组的 子序列 是从原数组中删除一些元素后（也可能一个也不删除）剩余元素不改变顺序组成的数组。比方说，[2,7,4] 是 [4,2,3,7,2,1,4] 的一个子序列（加粗元素），但是 [2,4,2] 不是。

示例 1：
输入：nums = [4,2,5,3]
输出：7
解释：最优子序列为 [4,2,5] ，交替和为 (4 + 5) - 2 = 7 。
示例 2：

输入：nums = [5,6,7,8]
输出：8
解释：最优子序列为 [8] ，交替和为 8 。
示例 3：

输入：nums = [6,2,1,2,4,5]
输出：10
解释：最优子序列为 [6,1,5] ，交替和为 (6 + 5) - 1 = 10 。

提示：

1 <= nums.length <= 105
1 <= nums[i] <= 105

方法：想了一会，没什么思路。
如果枚举所有的子序列，然后再对每一个子序列求交易和，最后更新最大值。
枚举比较容易，每个数字都有两种情况，选择放入当前子序列，或者不放入。

maxAlternatingSum1这个函数可以求解，但是肯定很慢。
这个写法也不能用记忆化。如何改造？

方法：贪心算法 可以转化为买卖股票问题(买卖多次，每次只能持有一股)
把数组中的每个数认为是股票的价格。由于本题将奇数位置认为是卖出,
从右到左对数组将相邻两天进行做差，然后累计所有的正数。
[6,2,1,2,4,5]->[4,1,-1,-2,-1,5] -> 10
*/
func maxAlternatingSum1(nums []int) int64 {
	var ans int64 = math.MinInt64
	// flag为正时，表示+  否则为-
	var dfs func(i int, sum int64, flag bool)
	dfs = func(i int, sum int64, flag bool) {
		if i >= len(nums) {
			if sum > ans {
				ans = sum
			}
			return
		}
		// 选择sum[i]
		if flag {
			dfs(i+1, sum+int64(nums[i]), !flag)
		} else {
			dfs(i+1, sum-int64(nums[i]), !flag)
		}
		// 不选择sum[i]
		dfs(i+1, sum, flag)
	}
	dfs(0, 0, true)
	return ans
}

/*
买卖股票问题的解法
*/
func maxAlternatingSum2(nums []int) int64 {
	var sum int64
	diff := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			diff[i] = nums[i]
		} else {
			diff[i] = nums[i] - nums[i-1]
		}
		if diff[i] > 0 {
			sum += int64(diff[i])
		}
	}
	return sum
}

/*
改造maxAlternatingSum1这个递归函数
使之能够进行记忆化搜索
思考之间的dfs中的flag,其实是想知道，如果选择了当前元素，应该是加上它还是减掉它。
而加减的核心在于子序列当前的长度是奇数还是偶数。
如果当前长度是偶数，则当前元素在奇数位置上，应该sum + nums[i]。长度为奇数的同理。

这样思考，从左到右考虑nums中的每一个元素。如果知道了nums[0...i-1]这段子数组的最大交替和，
考虑nums[i]这个元素：
 1. 选择nums[i] -> 如果之前最大交替和子序列长度是奇数，dfs(i) = dfs(i-1) -nums[i]
    如果之前最大交替和子序列长度是偶数，dfs(i) = dfs(i-1) + nums[i]
 2. 不选nums[i] -> dfs(i) = dfs(i-1)

因为，dfs函数应该有两个参数（i , flag）一个返回值，返回nums[0...i]的最大交替和
flag为true时，表示最大交替子序列长度为奇数。

递归边界：只有一个元素。长度只能是奇数。（i==0 && flag ==true）-> dfs(0 , true) = nums[0]
只有一个元素，长度为偶数。那么长度只能为0(这个时候不能选）。 给dfs(0.false) = 0
*/
func maxAlternatingSum3(nums []int) int64 {

	var dfs func(i int, flag bool) int
	dfs = func(i int, flag bool) int {
		if i == 0 {
			if flag {
				return nums[0]
			}
			return 0
		}
		// 选择nums[i]
		var m1, m2 int
		if flag {
			m1 = dfs(i-1, !flag) + nums[i]
		} else {
			m2 = dfs(i-1, !flag) - nums[i]
		}
		// 不选择
		// 三者取最大
		return max(m1, m2, dfs(i-1, flag))
	}
	return int64(dfs(len(nums)-1, true))

}

/*
1:1翻译成动态规划
dfs有两个参数，所以dp数组应该有二维。
为了方便在二维数组中表示i和flag，将i和flag统一类型。
使用flag->0表示偶数长度，使用flag->1表示奇数长度。
*/

func maxAlternatingSum4(nums []int) int64 {

	dp := make([][2]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = [2]int{}
	}
	dp[0][1] = nums[0]
	dp[0][0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+nums[i])
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-nums[i])
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return int64(dp[len(dp)-1][1])
}

func TestMaxAlternatingSum(t *testing.T) {
	//fmt.Println(maxAlternatingSum1([]int{4, 2, 5, 3}))
	//fmt.Println(maxAlternatingSum3([]int{4, 2, 5, 3}))
	//fmt.Println(maxAlternatingSum1([]int{6, 2, 1, 2, 4, 5}))
	//fmt.Println(maxAlternatingSum3([]int{6, 2, 1, 2, 4, 5}))
	//fmt.Println(maxAlternatingSum4([]int{6, 2, 1, 2, 4, 5}))
	fmt.Println(maxAlternatingSum4([]int{5, 6, 7, 8}))
	fmt.Println(maxAlternatingSum3([]int{5, 6, 7, 8}))
}
