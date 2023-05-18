package _0_dp

/*
nums[0] = 0
nums[1] = 1
当 2 <= 2 * i <= n 时，nums[2 * i] = nums[i]
当 2 <= 2 * i + 1 <= n 时，nums[2 * i + 1] = nums[i] + nums[i + 1]
返回生成数组 nums 中的 最大 值。
*/
func getMaximumGenerated(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	var ans int
	for j := 2; j <= n; j++ {
		if j%2 == 0 {
			arr[j] = arr[j/2]
		} else {
			arr[j] = arr[j/2] + arr[j/2+1]
		}
		ans = max(ans, arr[j])
	}
	return ans
}
