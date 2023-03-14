package mid

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/longest-consecutive-sequence/?favorite=2cktkvj

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
时间复杂度要求O(n)

没什么思路： 如果使用暴力解法应该怎么解？

hash解法：枚举
遍历数组，当检查到下标i的位置的时候，从nums[i]开始进行枚举 nums[i]+x是否在数组中存在
x=[1,2,3,4...]

*/

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	mp := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = struct{}{}
	}
	longest := 1

	for i := 0; i < len(nums); i++ {
		value := nums[i]
		l := 1
		if _, found := mp[value-1]; found { //说明以当前数字为起点的子序列已经被搜索过
			continue
		}

		for {
			if _, found := mp[value+1]; found {
				l++
				if l > longest {
					longest = l
				}
				value++
			} else {
				break
			}
		}
	}

	return longest

}

func TestLongestConsecutive(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{}))
	fmt.Println(longestConsecutive([]int{0}))
}
