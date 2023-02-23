package _3_hashtable

/*
代码中的解法体现了一种思想，就是只需要回答限定的问题。
如果能够搜索出全部的组合，那么肯定是能够回答有多少种组合。
但是能够回到有多少种组合，却并不一定能够回答有哪些组合。

*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	count := 0

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			m[num1+num2]++ // 所有组合计算一次，并统计频数
		}
	}

	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			count += m[-num3-num4] // 同样的模式找匹配
		}
	}
	return count
}
