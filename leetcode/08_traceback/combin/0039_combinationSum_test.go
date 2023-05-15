package combin

import (
	"fmt"
	"sort"
	"testing"
)

/*
39. 组合总和
https://leetcode.cn/problems/combination-sum/
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
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
