package other

import "sort"

/*
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

排序法，需要考虑2种情况：
1. 缺失的数字在[0...n-1]之间 这之间缺失的数字，排序后，nums[i]!=i。。例如，[3，0，1]这个数组
2. 缺失的数字就是n。排序后的数组满足nums[i]==i。但是肯定缺了数字，那么这个数字必定是n。例如，[0,1]这个数组。
*/
func missingNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n
}
