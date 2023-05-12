package _5_binarySearch

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

方法：二分查找

	x>= target

需要找两个下标：从左到右第一个等于target的数字的下标、从左到右第一个大于target的数字的下标-1
*/
func searchRange1(nums []int, target int) []int {
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

/*


查找nums[i]==target的最后一个，相当于查找nums[i]==target+1的第一个的索引前面一个。
如果能查找到的话。所以，两个位置的查找可以用一个函数搞定。

*/

func searchRange2(nums []int, target int) []int {
	var ans = []int{-1, -1}
	start := binaryLowBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return ans
	}

	end := binaryLowBound(nums, target+1) - 1
	ans[0] = start
	ans[1] = end

	return ans
}
func binaryLowBound(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if mid-1 >= 0 && nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		}
	}
	return left
}

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange1([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange1([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange1([]int{2, 2}, 3))
	fmt.Println(searchRange2([]int{2, 2}, 3))

}
