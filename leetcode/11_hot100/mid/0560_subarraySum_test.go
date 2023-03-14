package mid

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/subarray-sum-equals-k/?favorite=2cktkvj

给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

输入：nums = [1,1,1], k = 2
输出：2

解题思路：
1. 暴力解法 会出现一定程度的超时
(排序后使用滑动窗口 md这个题不能用滑动窗口做)
2.前缀数组



*/

func subarraySum1(nums []int, k int) int {

	// 统计从[i...j]这个区间的和是否等于k
	var sum func(arr []int, beginIndex int, endIndex int) int
	sum = func(arr []int, beginIndex int, endIndex int) int {
		sum := 0
		for i := beginIndex; i <= endIndex; i++ {
			sum += arr[i]
		}
		return sum
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if sum(nums, i, j) == k {
				res++
			}
		}
	}
	return res
}

/*
前缀数组
[j...i] 这个段区间的和等于 prefix[i] - prefix[j-1] (这里注意是j-1)
这个题有点两数之和，使用hashmap进行速度优化的感觉。
参考题解：https://leetcode.cn/problems/subarray-sum-equals-k/solutions/238572/he-wei-kde-zi-shu-zu-by-leetcode-solution/
*/

func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	mp := make(map[int]int)
	mp[0] = 1

	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if c, found := mp[pre-k]; found {
			count += c
		}
		if _, found := mp[pre]; found {
			mp[pre]++
		} else {
			mp[pre] = 1
		}
	}
	return count
}

func TestSubArray(t *testing.T) {

	fmt.Println(subarraySum2([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum2([]int{1, 1, 1}, 2))
}
