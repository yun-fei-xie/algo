package _5_doublePointer

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/4sum/description/
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

方法：对向双指针，降维到2数之和。
需要去重，对[i,j,left,right]的去重条件需要注意。


*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	ans := make([][]int, 0)

	for i := 0; i < length-3; i++ {
		// 对i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 对j去重 为什么是j>i+1
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 两数之和
			for left, right := j+1, length-1; left < right; {

				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum > target {
					right--
				} else if sum < target {
					left++
				} else { //sum == target
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					// 对left 去重 注意边界 防止数组越界
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					// 对right 去重
					for left < right && nums[right] == nums[right+1] {
						right--
					}

				}

			}

		}
	}
	return ans
}

func TestFourSum(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
	// [-2,-1,-1,1,1,2,2]

}
