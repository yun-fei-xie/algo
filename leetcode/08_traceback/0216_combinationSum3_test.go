package _8_traceback

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/combination-sum-iii/description/
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
*/
func combinationSum3(k int, n int) [][]int {
	var res = [][]int{}
	var path = []int{}

	var dfs func(startIndex int, sum int)
	dfs = func(startIndex int, sum int) {
		if len(path) == k {
			if sum == n {
				var temp = make([]int, len(path)) // 必须显示申明长度，不然没法copy
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}

		//if sum > n { // 剪枝的条件之一  放在for循环里面，如果满足这个条件，不需要进入递归
		//	return
		//}

		for i := startIndex; i <= 9; i++ {
			if sum+i > n { // 剪枝操作
				break
			}
			path = append(path, i)
			dfs(i+1, sum+i)
			path = path[:len(path)-1]
		}

	}
	dfs(1, 0)
	return res
}

func TestCombinationSum3(t *testing.T) {
	k, n := 3, 9
	res := combinationSum3(k, n)
	fmt.Println(res)
}
