package mid

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/description/
解法1：排序后使用双指针
解法2：
*/

func findUnsortedSubarray(nums []int) int {
	left := 0
	right := len(nums) - 1
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	sort.Ints(nums)

	for left <= right && nums[left] == numsCopy[left] {
		left++
	}

	for left <= right && nums[right] == numsCopy[right] {
		right--
	}
	return right - left + 1
}

/*
优化空间复杂度 不使用排序
分情况讨论：
*/
func findUnsortedSubarray2(nums []int) int {
	return -1
}

func TestFindUnsortedSubarray(t *testing.T) {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 3, 3}))
}
