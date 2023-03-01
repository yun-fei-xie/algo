package _9_greedy

/*
https://leetcode.cn/problems/jump-game/description/
只要能够覆盖 就一定能调到终点 有点类似于转化思想
https://programmercarl.com/0055.%E8%B7%B3%E8%B7%83%E6%B8%B8%E6%88%8F.html#%E6%80%9D%E8%B7%AF
*/
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	max := func(x int, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	cover := 0 // 能够覆盖的范围
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
