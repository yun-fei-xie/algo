package subset

import (
	"fmt"
	"testing"
)

/*
78. 子集
https://leetcode.cn/problems/subsets/

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
长度为0的、长度为1的、长度为len(nums)-1的
*/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	path := make([]int, 0)

	var dfs func(nums []int, startIndex int, depth int, targetLength int)
	dfs = func(nums []int, startIndex int, depth int, targetLength int) {
		if depth == targetLength {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(nums, i+1, depth+1, targetLength)
			path = path[:len(path)-1]
		}
	}
	// 在深度优先遍历中使用for循环  找出每个长度
	for d := 1; d <= len(nums); d++ {
		dfs(nums, 0, 0, d)
	}
	return res
}

/*
每一个字符都有两种选择
1. 被放入当前集合
2. 不放入
*/
func subsets2(nums []int) [][]int {
	var ans = make([][]int, 0)
	var length = len(nums)
	var path = make([]int, 0)
	var dfs func(startIndex int)

	dfs = func(startIndex int) {
		if startIndex == length {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		// 每个元素可以选 也可以不选
		path = append(path, nums[startIndex])
		dfs(startIndex + 1)

		// 不选 不选的话，就需要把刚刚添加到path中的值去掉
		path = path[0 : len(path)-1]
		dfs(startIndex + 1)
	}

	dfs(0)
	return ans
}

func TestSubSet(t *testing.T) {
	fmt.Println(subsets2([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1, 2, 3}))
}
