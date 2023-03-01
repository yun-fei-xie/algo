package _8_traceback

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/subsets-ii/description/
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
