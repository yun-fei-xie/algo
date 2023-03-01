package _8_traceback

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/permutations-ii/

输入：nums = [1,1,2]
输出：
[[1,1,2],

	[1,2,1],
	[2,1,1]]
*/
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]int, 30) // [-10 ,10]
	targetLength := len(nums)

	var dfs func(nums []int, depth int)
	dfs = func(nums []int, depth int) {
		if depth == targetLength {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		mp := make(map[int]struct{}) // 记录同一层是否重复 (同一个父节点下是否重复)

		for i := 0; i < len(nums); i++ {
			if _, found := mp[nums[i]]; found {
				continue
			}

			if used[i] == 0 {
				path = append(path, nums[i])
				used[i] = 1
				dfs(nums, depth+1)
				path = path[:len(path)-1]
				used[i] = 0

				mp[nums[i]] = struct{}{}
			}
		}
	}
	dfs(nums, 0)
	return res
}

func TestPermuteUnique(t *testing.T) {

	nums := []int{1, 1, 2}
	res := permuteUnique(nums)
	fmt.Println(res)

}
