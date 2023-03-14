package mid

import (
	"fmt"
	"testing"
)

/*
34. 在排序数组中查找元素的第一个和最后一个位置
https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?favorite=2cktkvj
中等
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

本题考察的是二分查找的变种
解法：
1. 直接遍历，算法复杂度为O(n)，不符合题意，因此需要使用到二分查找（二分查找的变种）

需要找两个下标：从左到右第一个等于target的数字的下标、从左到右第一个大于target的数字的下标-1
*/
func searchRange(nums []int, target int) []int {
	begin := binanryFirst(nums, 0, len(nums)-1, target)
	end := binaryLatest(nums, 0, len(nums)-1, target)
	return []int{begin, end}
}

func binanryFirst(nums []int, left int, right, target int) int {

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target { // 查找最左边第一个等于target的元素
				return mid
			} else {
				right = mid - 1 // 继续向左查询
			}
		}
	}
	return -1
}

func binaryLatest(nums []int, left int, right int, target int) int {
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {

			if mid == len(nums)-1 || nums[mid+1] != target { // 查找最后一个等于target的元素
				return mid
			} else {
				left = mid + 1 // 继续向右去查询
			}

		}
	}
	return -1
}

func TestSearchRange(t *testing.T) {
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{2, 2}, 3))

}
