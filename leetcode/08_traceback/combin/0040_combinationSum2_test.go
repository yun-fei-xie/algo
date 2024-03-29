package combin

import (
	"fmt"
	"sort"
	"testing"
)

/*
40. 组合总和 II
https://leetcode.cn/problems/combination-sum-ii/description/

给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。

这个题结果容易重复，是因为数组中本身有重复的元素,但是又要求不能有重复的"组合" [2,1,2]和[2,2,1]重复
树层重复 和 树枝重复 。从根节点一路下来是数枝）
[10,1,2,7,6,1,5]-> [1,1,2,5,6,7,10]
*/
func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	var used = make([]bool, len(candidates))
	var ans [][]int
	var path []int

	var traceback func(startIndex int, sum int)
	traceback = func(startIndex int, sum int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := startIndex; i < len(candidates) && candidates[i]+sum <= target; i++ {
			if i != 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
				continue
			}
			path = append(path, candidates[i])
			used[i] = true
			traceback(i+1, sum+candidates[i])
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	traceback(0, 0)
	return ans

}

func TestCombinationSum2(t *testing.T) {
	arr := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	res := combinationSum2(arr, target)
	fmt.Println(res)

}
