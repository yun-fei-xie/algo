package _5_binarySearch

import (
	"fmt"
	"testing"
)

/*
162. 寻找峰值
https://leetcode.cn/problems/find-peak-element/

峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
*/
func findPeakElement(nums []int) int {
	var length = len(nums)
	var left = 0
	var right = length - 1
	// 循环结束 left == right
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func TestFindPeakElement(t *testing.T) {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
