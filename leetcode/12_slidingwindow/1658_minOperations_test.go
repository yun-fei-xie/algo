package _2_slidingwindow

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/description/

给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。
如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。
输入：nums = [1,1,4,2,3], x = 5
输出：2
解释：最佳解决方案是移除后两个元素，将 x 减到 0 。

可以看到，这个问题是从两端抽取连续的元素。如果用双指针从两端开始考虑，确实比较难。
但是可以换个思路。题目要求找两端连续的子数组和等于x，
就等价于求子数组和等于 sum(nums[0...len(nums)])-x
符合条件的子数组的最大长度。

这样就转换为一个滑动窗口的问题（或者叫双指针 这个窗口的大小会改变）
*/

func minOperations(nums []int, x int) int {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	target := totalSum - x

	var sum int
	var ans int = -1
	for left, right := 0, 0; right < len(nums); right++ {
		sum += nums[right]
		// target可能是负数，sum最小是0，如果不加left<=right这个限制条件，数组可能会越界。
		for sum > target && left <= right {
			sum -= nums[left]
			left++
		}
		if sum == target {
			ans = max(ans, right-left+1)
		}
		// if sum<target 直接下一轮循环 所以这里不用写
	}
	if ans == -1 {
		return -1
	} else {
		return len(nums) - ans
	}
}

func TestMinOperations(t *testing.T) {
	//fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	//fmt.Println(minOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(minOperations([]int{1, 1}, 3))
}
