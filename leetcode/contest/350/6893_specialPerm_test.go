package _50

import (
	"fmt"
	"testing"
)

/*
6893. 特别的排列
https://leetcode.cn/problems/special-permutations/

如何从数组中生成所有的排列 写回溯肯定会超时

可以这样想，把全排列看做是位置，现在从nums中选数字填位置。然后用排列组合计算结果。
哪里会重复，假如前面3位是这样选的
4,1,2,.....
2,1,4,.....
后面的选择就会重复。
*/
func specialPerm(nums []int) int {
	var ans int
	mod := 1000000000 + 7
	path := make([]int, 0)
	used := make([]bool, len(nums))
	length := len(nums)

	var traceback func(depth int)
	traceback = func(depth int) {
		if depth == length {
			ans = (ans + 1) % mod
			return
		}

		for i := 0; i < length; i++ {
			if used[i] == false {
				if len(path) == 0 || (nums[i]%path[len(path)-1] == 0) || (path[len(path)-1]%nums[i] == 0) {
					used[i] = true
					path = append(path, nums[i])
					traceback(depth + 1)
					used[i] = false
					path = path[0 : len(path)-1]
				}
			}
		}
	}
	traceback(0)
	return ans
}

/*
 */

func specialPerm2(nums []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(nums)
	m := 1 << n
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	// i就是状态（联系集合与位运算）
	// j表示上一个被选择的数(因此，需要初始时选一个数字，体现在最后的for循环中)
	dfs = func(i, j int) (res int) {
		// i==0 表示所有的位置都选上了。
		if i == 0 {
			return 1 // 找到一个特别排列
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		//这里的i就是bool数组的整数形式 如何确认当前nums中的第k位（k从0开始）是否没有被选择过（选过就是0 ，没选就是1）
		//将i左移k位然后与1进行与运算。
		//将k这个位置放到bool数组中，将1右移动k位，然后和i进行或运算
		for k, x := range nums {
			if i>>k&1 == 1 && (nums[j]%x == 0 || x%nums[j] == 0) {
				res = (res + dfs(i^(1<<k), k)) % mod
			}
		}
		*p = res
		return
	}
	// 选定第一个数字
	for j := range nums {
		ans = (ans + dfs((m-1)^(1<<j), j)) % mod
	}
	return
}

func TestPermute(t *testing.T) {
	fmt.Println(specialPerm([]int{2, 3, 6}))
	fmt.Println(specialPerm([]int{1, 4, 3}))
}
