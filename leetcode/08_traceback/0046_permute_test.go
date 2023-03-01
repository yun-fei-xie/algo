package _8_traceback

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/permutations/
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

每次都重新0开始遍历，找出一个没有用过的数字加入到path中 （startIndex 这个参数可以不要 0047扔掉了这个参数）
*/

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]int, len(nums))
	targetLength := len(nums)

	var dfs func(nums []int, startIndex int, depth int)
	dfs = func(nums []int, startIndex int, depth int) {
		if depth == targetLength {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] == 0 {
				path = append(path, nums[i])
				used[i] = 1
				dfs(nums, 0, depth+1)
				used[i] = 0
				path = path[:len(path)-1]
			}
		}
	}
	dfs(nums, 0, 0)
	return res

}

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)

}
