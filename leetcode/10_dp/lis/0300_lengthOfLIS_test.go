package lis

import (
	"fmt"
	"math"
	"testing"
)

/*
300.最长递增子序列
https://leetcode.cn/problems/longest-increasing-subsequence/?envType=study-plan-v2&envId=dynamic-programming

给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
示例 1：
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。


*/

/*
递归解法：前序遍历。每个数字都可以放在序列后面，或者不放在序列后面。
这个时间复杂度是2^n。（每一个元素都有两种情况）
如果要放在序列后面就要满足条件（比序列中的最后一个元素大）
递归结束时，比较序列的长度
[2,3,7,101]
前序遍历没办法用记忆化搜索。思考一下，如何倒着向前推导。
*/
func lengthOfLIS(nums []int) int {
	var ans int = math.MinInt32
	var dfs func(nums []int, startIndex int, preNum int, sequenceLen int)
	dfs = func(nums []int, startIndex int, preNum int, sequenceLen int) {
		if startIndex >= len(nums) { // 枚举完毕
			if sequenceLen > ans {
				ans = sequenceLen
			}
			return
		}

		// 放进去
		if nums[startIndex] > preNum {
			dfs(nums, startIndex+1, nums[startIndex], sequenceLen+1)
		}
		// 不放进去
		dfs(nums, startIndex+1, preNum, sequenceLen)

		// 在这里改造：放进去的长度是sequenceLen+1 , 不放进去的长度是sequence

	}

	dfs(nums, 0, math.MinInt32, 0)

	return ans
}

/*
记忆化搜索改造。使用记忆化搜索应该具备后序的特征，例如，斐波那契数列的计算。
后续遍历改造
*/
func lengthOfLIS2(nums []int) int {
	var ans int = 1
	//递归函数的定义，返回[startIndex,len(nums))的最长递增子序列的长度 nums[startIndex]必须在序列中
	var dfs func(startIndex int) int
	dfs = func(startIndex int) int {
		if startIndex >= len(nums)-1 {
			return 1
		}

		var ret = 1
		for i := startIndex + 1; i < len(nums); i++ {
			//尝试将nums[startIndex]放在[startIndex+1,len(nums)-1] 中的各个子序列的前面
			d := dfs(i)
			if nums[startIndex] < nums[i] {
				d++
			}
			ret = max(ret, d)
		}
		//	fmt.Printf("startIndex: %d  num: %d  ret %d\n", startIndex, nums[startIndex], ret)
		ans = max(ans, ret)
		return ret
	}
	dfs(0)
	return ans
}

func lengthOfLIS3(nums []int) int {
	var ans = 1
	mem := make([]int, len(nums))
	//递归函数的定义，返回[startIndex,len(nums))的最长递增子序列的长度 nums[startIndex]必须在序列中
	var dfs func(nums []int, startIndex int) int
	dfs = func(nums []int, startIndex int) int {
		if startIndex >= len(nums)-1 {
			mem[startIndex] = 1
			return mem[startIndex]
		}

		var ret = 1
		for i := startIndex + 1; i < len(nums); i++ {
			//尝试将nums[startIndex]放在[startIndex+1,len(nums)-1] 中的各个子序列的前面
			if mem[i] == 0 {
				mem[i] = dfs(nums, i)
			}
			if nums[startIndex] < nums[i] {
				ret = max(ret, mem[i]+1)
			}
		}
		mem[startIndex] = ret
		// fmt.Printf("startIndex: %d  num: %d  ret %d\n", startIndex, nums[startIndex], ret)
		ans = max(ans, mem[startIndex])
		return mem[startIndex]
	}
	dfs(nums, 0)
	return ans
}

/*
动态规划
*/
func lengthOfLIS4(nums []int) int {
	mem := make([]int, len(nums))
	for i := 0; i < len(mem); i++ {
		mem[i] = 1
	}
	var max func(int, int) int
	max = func(i1 int, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	var ans = math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		//搜索以i为起点的最长递增子序列
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] { // 可以拼接到后面
				mem[i] = max(mem[i], mem[j]+1)
			}
		}
		ans = max(mem[i], ans)
	}

	return ans
}

func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

func TestLengthOfLIS(t *testing.T) {
	//fmt.Println(lengthOfLIS3([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS3([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS3([]int{4, 10, 4, 3, 8, 9}))
	fmt.Println(lengthOfLIS3([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS2([]int{0, 1, 0, 3, 2, 3}))
	// 0, 1, 0, 3, 2, 3
	// 4, 3, 3, 1 ,2 ,1
	// 4, 3, 0, 1, 2, 1
}
