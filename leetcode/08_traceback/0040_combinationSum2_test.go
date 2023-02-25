package _8_traceback

import (
	"fmt"
	"testing"
)

/*
这个题结果容易重复，是因为数组中本身有重复的元素
*/
func combinationSum2(candidates []int, target int) [][]int {

	var res = [][]int{}
	var path = []int{}

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
			dfs(arr, i+1, sum+arr[i])
			path = path[:len(path)-1]
		}
	}
	dfs(candidates, 0, 0)
	return res
}

func TestCombinationSum2(t *testing.T) {
	arr := []int{2, 5, 2, 1, 2}
	target := 5

	res := combinationSum2(arr, target)
	fmt.Println(res)

}
