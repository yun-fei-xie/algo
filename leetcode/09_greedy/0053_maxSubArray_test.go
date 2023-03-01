package _9_greedy

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/maximum-subarray/
暴力解法 二重循环 会超时
*/
func maxSubArrayTimeOut(nums []int) int {
	max := math.MinInt

	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			if count > max {
				max = count
			}
		}
	}
	return max

}

/*
贪心算法
局部最优：当前“连续和”为负数的时候立刻放弃，从下一个元素重新计算“连续和”，因为负数加上下一个元素 “连续和”只会越来越小。
全局最优：选取最大“连续和”
*/

func maxSubArray(nums []int) int {
	max := math.MinInt
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > max { //先记录最值 然后再判断是否小于0
			max = count
		}
		if count < 0 {
			count = 0
		}
	}
	return max
}

func TestMaxSubArray(t *testing.T) {

	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArrayTimeOut(arr)
	fmt.Println(res)

}
