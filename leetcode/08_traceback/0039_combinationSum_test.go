package _8_traceback

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/combination-sum/
多画树形图

这一题的输入数组是没有重复的
40题是有重复的数组
*/
func combinationSum(candidates []int, target int) [][]int {

	var res = [][]int{}
	var path = []int{}
	// 递归终止 sum > target

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
			path = append(path, arr[i])
			dfs(arr, i, sum+arr[i])
			path = path[:len(path)-1]
		}
	}
	dfs(candidates, 0, 0)
	return res
}

func TestCombinationSum(t *testing.T) {
	//candidates := []int{2, 3, 6, 7}
	//target := 7
	//res := combinationSum(candidates, target)
	//fmt.Println(res)

	candidates2 := []int{2, 3, 5}
	target2 := 8
	res2 := combinationSum(candidates2, target2)
	fmt.Println(res2)

}
