package _49

import "sort"

/*
只需要看前三个数字
去掉最大最小
*/
func findNonMinOrMax(nums []int) int {

	sort.Ints(nums)
	if len(nums) <= 2 {
		return -1
	}
	return nums[1]
}
