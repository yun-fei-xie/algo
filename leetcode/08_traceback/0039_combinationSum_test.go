package _8_traceback

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/combination-sum/
多画树形图

这一题的输入数组是没有重复的
40题是有重复的数组
*/
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	sort.Ints(candidates) // 排序
	var traceback func(startIndex int, sum int)
	traceback = func(startIndex int, sum int) {
		if sum >= target {
			if sum == target {
				temp := make([]int, len(path))
				copy(temp, path)
				ans = append(ans, temp)
			}
			return
		}

		for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {

			path = append(path, candidates[i])
			traceback(i, sum+candidates[i])
			path = path[:len(path)-1]
		}

	}
	traceback(0, 0)
	return ans
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
