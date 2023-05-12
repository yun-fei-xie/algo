package _5_doublePointer

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/3sum/description/

给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

方法：外层枚举，内层使用0167题的方法（对向双指针）

*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	ans := make([][]int, 0)
	//规定一下三个数的顺序 i<j<k
	for i := 0; i < length-2; i++ {
		// 处理i重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, length-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--

				// 为什么只能在sum==0的情况下处理重复元素的问题。（放在外面处理还会造成数组下标越界）
				// 处理j重复 要使用j<k限制，不然数组可能会越界 [0,0,0]这种情况
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				// 处理k重复 要使用j<k限制，不然数组可能会越界 [0,0,0]这种情况
				for j < k && nums[k] == nums[k+1] {
					k--
				}

			} else if sum > 0 {
				k--
			} else { // sum < 0
				j++
			}

		}
	}
	return ans
}

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}
