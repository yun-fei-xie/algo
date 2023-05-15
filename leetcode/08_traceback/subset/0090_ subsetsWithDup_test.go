package subset

import (
	"fmt"
	"sort"
	"testing"
)

/*
90. 子集 II
https://leetcode.cn/problems/subsets-ii/description/
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

示例 1：
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
示例 2：

输入：nums = [0]
输出：[[],[0]]
*/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	res = append(res, []int{})
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

		for i := startIndex; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 { // 树层去重
				continue
			}
			path = append(path, nums[i])
			used[i] = 1
			dfs(nums, i+1, depth+1, targetLength)
			used[i] = 0
			path = path[:len(path)-1]
		}
	}

	for d := 1; d <= len(nums); d++ {
		dfs(nums, 0, 0, d)
		used = make([]int, len(nums))
	}
	return res
}

func TestSubSetWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}
