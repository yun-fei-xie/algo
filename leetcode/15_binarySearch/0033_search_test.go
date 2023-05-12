package _5_binarySearch

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：
输入：nums = [1], target = 0
输出：-1
*/

/*
方法一：使用两次二分
第一次二分找出数组中的最小值所在的位置，通过最小值将数组分为两段有序的子数组。
通过与第二段的上界与第一段的下界进行比较，确定第二次使用二分法的区间。
nums = [4,5,6,7,0,1,2], target = 3
第一次二分可以知道，0所在的位置4，将数组分为2段arr1->[4,5,6,7]和arr2->[0,1,2]
target>max(arr2) 所以需要在arr1中使用普通的二分查找。


*/

func search(nums []int, target int) int {
	minIndex := findMinIndex(nums)
	if target > nums[len(nums)-1] {
		return binarySearch(nums, 0, minIndex-1, target)
	} else {
		return binarySearch(nums, minIndex, len(nums)-1, target)
	}

}

func binarySearch(nums []int, left int, right int, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func findMinIndex(nums []int) int {
	var length = len(nums)
	var left = 0
	var right = length - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
}
