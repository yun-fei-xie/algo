package _8_traceback

import (
	"fmt"
	"testing"
)

/*
491. 递增子序列

https://leetcode.cn/problems/non-decreasing-subsequences/
输入：nums = [4,6,7,7]
输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]

注意不能排序，要在原数组已有的顺序上寻找子序列
但是不排序如何去重。。。

看题解后，知道了如何在函数中记录树种的同一层是否使用过
*/
func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]int, len(nums))

	var dfs func(nums []int, startIndex int, depth int, targetLength int)
	dfs = func(nums []int, startIndex int, depth int, targetLength int) {
		if depth == targetLength {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		mp := make(map[int]bool, 0) //只在树的单层中使用
		for i := startIndex; i < len(nums); i++ {

			// 同层去重
			if _, found := mp[nums[i]]; found {
				continue
			} else {
				mp[nums[i]] = true
			}

			// 限制条件2 递增序列
			if len(path) != 0 {
				last := path[len(path)-1]
				if nums[i] < last {
					continue
				}
			}

			path = append(path, nums[i])
			used[i] = 1
			dfs(nums, i+1, depth+1, targetLength)
			used[i] = 0
			path = path[:len(path)-1]
		}
	}

	for d := 2; d <= len(nums); d++ { // 最少2个元素
		dfs(nums, 0, 0, d)
		used = make([]int, len(nums))
	}
	return res
}

func TestSubSeq(t *testing.T) {
	nums1 := []int{4, 6, 7, 7}
	res1 := findSubsequences(nums1)
	fmt.Println(res1)

	nums2 := []int{4, 4, 3, 2, 1}
	res2 := findSubsequences(nums2)
	fmt.Println(res2)
}
