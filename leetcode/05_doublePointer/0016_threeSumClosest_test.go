package _5_doublePointer

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/3sum-closest/description/
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。
假定每组输入只存在恰好一个解。
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

方法：既然是求最接近的值（做差求绝对值）
具体的思路和三数之和一样。内层使用对向双指针。
需要注意的是，这一题不需要去重，题目明确说明没有重复。
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var length = len(nums)
	var ans int
	var diff int = math.MaxInt

	for i := 0; i < length-2; i++ {

		for left, right := i+1, length-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if sum > target {
				newDiff := int(math.Abs(float64(sum - target)))
				if newDiff < diff {
					diff = newDiff
					ans = sum
				}
				right--
			} else if sum < target {
				newDiff := int(math.Abs(float64(sum - target)))
				if newDiff < diff {
					diff = newDiff
					ans = sum
				}
				left++
			} else { // sum == target
				return sum
			}
		}
	}
	return ans
}

func TestThreeSumClosest(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
