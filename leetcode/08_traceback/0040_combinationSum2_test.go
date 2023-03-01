package _8_traceback

import (
	"fmt"
	"sort"
	"testing"
)

/*
这个题结果容易重复，是因为数组中本身有重复的元素,但是又要求不能有重复的"组合" [2,1,2]和[2,2,1]重复
树层重复 和 树枝重复 。从根节点一路下来是数枝）
https://leetcode.cn/problems/combination-sum-ii/description/
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 必须先排序 保证
	var res = [][]int{}
	var path = []int{}
	var used = make([]int, len(candidates))

	var dfs func(arr []int, startIndex int, sum int)
	dfs = func(arr []int, startIndex int, sum int) {
		if sum >= target {
			if sum == target {
				temp := make([]int, len(path))
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}

		for i := startIndex; i < len(arr); i++ {

			if i > 0 && arr[i] == arr[i-1] && used[i-1] == 0 {
				continue
			}

			path = append(path, arr[i])

			used[i] = 1
			dfs(arr, i+1, sum+arr[i])
			used[i] = 0

			path = path[:len(path)-1]
		}
	}
	dfs(candidates, 0, 0)
	return res
}

func TestCombinationSum2(t *testing.T) {
	arr := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	res := combinationSum2(arr, target)
	fmt.Println(res)

}
