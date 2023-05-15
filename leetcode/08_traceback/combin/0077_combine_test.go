package combin

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/combinations/description/

给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

本质是一个树形问题
*/
func combine2(n int, k int) [][]int {

	var ans [][]int
	var path []int
	var traceBack func(startNumber int, depth int)
	traceBack = func(startNumber, depth int) {
		if depth == k {
			temp := make([]int, depth)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for i := startNumber; i <= n; i++ {
			if n-startNumber+1+depth < k {
				return
			}
			path = append(path, i)
			traceBack(i+1, depth+1)
			path = path[:len(path)-1]
		}
	}

	traceBack(1, 0)
	return ans
}

func combine(n int, k int) [][]int {
	var nums []int
	var res [][]int
	var path []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	var dfs func(nums []int, startIndex int)

	dfs = func(nums []int, startIndex int) {
		if len(path) == k { //控制递归的深度
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			//	dfs(nums, startIndex+1) // 重复在这个位置
			dfs(nums, i+1) // 这样写不会走回头路

			path = path[:len(path)-1]
		}

	}

	dfs(nums, 0)
	return res
}

func TestCombine(t *testing.T) {
	n, k := 4, 2
	res := combine(n, k)
	res2 := combine2(n, k)
	fmt.Println(res)
	fmt.Println(res2)

}
