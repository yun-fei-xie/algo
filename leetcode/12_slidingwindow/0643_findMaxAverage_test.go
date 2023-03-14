package _2_slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/maximum-average-subarray-i/
最简单的滑动窗口题目，维护固定大小的窗口
*/
func findMaxAverage(nums []int, k int) float64 {
	left := 0
	right := 0 + k - 1
	maxSum := math.MinInt64
	sum := 0
	for i := 0; i <= right; i++ {
		sum += nums[i]
	}
	maxSum = sum

	for right+1 < len(nums) {
		sum = sum + nums[right+1] - nums[left] // 滚动连续区间求和
		if sum > maxSum {
			maxSum = sum
		}
		// 窗口滑动
		left++
		right++
	}

	return float64(maxSum) / float64(k)

}

func TestFindMaxAverage(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	res := findMaxAverage(nums, 4)
	fmt.Println(res)

	nums2 := []int{0, 4, 0, 3, 2}
	res2 := findMaxAverage(nums2, 1)
	fmt.Println(res2)
}
