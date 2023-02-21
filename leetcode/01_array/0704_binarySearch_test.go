package _1_array

import (
	"fmt"
	"testing"
)

/**
https://leetcode.cn/problems/binary-search/
数组的二分查找
最重要的是定义边界条件 是大于还是大于等于这样的边界case
这就需要提前定义好search的范围，在哪个区间进行寻找
*/

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
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

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	res := search(nums, target)
	fmt.Println(res)
}
