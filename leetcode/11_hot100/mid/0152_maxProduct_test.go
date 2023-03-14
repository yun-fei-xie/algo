package mid

import "math"

/*
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
测试用例的答案是一个 32-位 整数。
子数组 是数组的连续子序列。

示例 1:
输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。


https://leetcode.cn/problems/maximum-product-subarray/?favorite=2cktkvj

第一思路：有点滑动窗口的感觉。试试？

product(i, j) = product(0, j) / product(0, i) 从数组 i 到 j 的累乘等于 从数组开头到 j 的累乘除以从数组开头到 i 的累乘(这里先忽略 0 的情况)，要考虑三种情况

累乘的乘积等于 0，就要重新开始
累乘的乘积小于 0，要找到前面最大的负数，这样才能保住从 i 到 j 最大
累乘的乘积大于 0，要找到前面最小的正数，同理！
参考题解：https://leetcode.cn/problems/maximum-product-subarray/solutions/17709/duo-chong-si-lu-qiu-jie-by-powcai-3/


*/

func maxProduct(nums []int) int {

	mul := 1
	max := math.MinInt64

	left := 0
	right := 0

	for left <= right && right < len(nums) {
		mul *= nums[right]
		if mul > max {
			max = mul
		}

		if mul == 0 {
			right++
			left = right
		} else if mul < 0 {

		}

	}

	return max
}

/*
暴力求解：二重循环枚举所有的子序列 -> 部分测试用例超时
*/

func maxProduct1(nums []int) int {

	var mul func(arr []int, begin int, end int) int
	mul = func(arr []int, begin int, end int) int {
		m := 1
		for i := begin; i <= end; i++ {
			m *= arr[i]
		}
		return m
	}

	max := math.MinInt64

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			r := mul(nums, i, j)
			if r > max {
				max = r
			}
		}
	}
	return max
}

func maxProduct2(nums []int) int {

	return 0
}
