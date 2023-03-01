package _8_traceback

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/subsets/

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

func TestSubSet(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)

}
