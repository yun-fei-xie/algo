package _348__test

import (
	"fmt"
	"math"
	"testing"
)

/*
2717. 半有序排列

https://leetcode.cn/problems/semi-ordered-permutation/

给你一个下标从 0 开始、长度为 n 的整数排列 nums 。
如果排列的第一个数字等于 1 且最后一个数字等于 n ，则称其为 半有序排列 。你可以执行多次下述操作，直到将 nums 变成一个 半有序排列 ：
选择 nums 中相邻的两个元素，然后交换它们。
返回使 nums 变成 半有序排列 所需的最小操作次数。
排列 是一个长度为 n 的整数序列，其中包含从 1 到 n 的每个数字恰好一次。

方法：模拟冒泡排序
*/
func semiOrderedPermutation(nums []int) int {
	length := len(nums)
	minIndex, min := -1, math.MaxInt32

	for i := 0; i < length; i++ {
		if nums[i] < min {
			min = nums[i]
			minIndex = i

		}
	}

	var ans int
	for j := minIndex; j > 0; j-- {

		if nums[j] < nums[j-1] {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	ans += minIndex

	maxIndex, max := -1, math.MinInt32
	for i := 0; i < length; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}

	ans += length - maxIndex - 1
	return ans

}

func TestSemiOrder(t *testing.T) {
	fmt.Println(semiOrderedPermutation([]int{2, 1, 4, 3}))
}
