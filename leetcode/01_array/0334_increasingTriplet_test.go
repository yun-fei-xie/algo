package _1_array

/*
https://leetcode.cn/problems/increasing-triplet-subsequence/
递增三元子序列

没有思路，就先暴力枚举，三重循环 会超时
*/
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	for i := 0; i <= len(nums)-3; i++ { // 第一个数字

		for j := i + 1; j <= len(nums)-2; j++ {

			for k := j + 1; k <= len(nums)-1; k++ {
				if nums[i] < nums[j] && nums[j] < nums[k] {
					return true
				}
			}
		}
	}
	return false
}

// 优化，nums = [2,1,5,0,4,6] 去掉不必要的循环。 比如，nums[i]=2 , nums[j]=1 那么直接可以continue了，不需要再做无用的检查
// 这一轮优化，可以再通过一些测试用例，但是依然不能AC。
func increasingTriplet2(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	for i := 0; i <= len(nums)-3; i++ { // 第一个数字

		for j := i + 1; j <= len(nums)-2; j++ {
			if nums[i] >= nums[j] {
				continue
			}
			for k := j + 1; k <= len(nums)-1; k++ {
				if nums[j] < nums[k] {
					return true
				}
			}
		}
	}
	return false
}
