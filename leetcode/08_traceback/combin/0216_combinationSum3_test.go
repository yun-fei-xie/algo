package combin

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
	var ans [][]int
	var path []int

	var traceBack func(startNumber int, depth int, sum int)
	traceBack = func(startNumber int, depth int, sum int) {
		//if depth > k || sum > n {
		//	return
		//}
		if depth == k && sum == n {
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for i := startNumber; i <= 9; i++ {
			if depth+1 > k || sum+i > n { // 剪枝
				continue
			}

			path = append(path, i)
			traceBack(i+1, depth+1, sum+i)
			path = path[:len(path)-1]
		}

	}
	traceBack(1, 0, 0)
	return ans
}

func TestCombinationSum3(t *testing.T) {
	k, n := 3, 9
	res := combinationSum3(k, n)
	fmt.Println(res)
}
