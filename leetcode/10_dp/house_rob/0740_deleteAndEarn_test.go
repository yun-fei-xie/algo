package house_rob

import (
	"fmt"
	"testing"
)

/*
740. 删除并获得点数
https://leetcode.cn/problems/delete-and-earn/description/
打家劫舍问题的变形问题
而打家劫舍问题和斐波那契问题在递推公式的结构上十分相似。

本题有两种场景
1.不包含重复元素
2.包含重复元素
使用下面这种基于基数排序的方法构造数组，可以将两个问题用同一种方法解决

解法参考lc官方题解
根据题意，在选择了元素x后，该元素以及所有等于 x−1 或 x+1 的元素会从数组中删去。
若还有多个值为x的元素，由于所有等于 x−1 或 x+1 的元素已经被删除，我们可以直接删除x并获得其点数。
因此若选择了x，所有等于x的元素也应一同被选择，以尽可能多地获得点数。
记元素x在数组中出现的次数为,我们可以用一个数组 sum 记录数组 nums\textit{nums}nums 中所有相同元素之和
*/
func deleteAndEarn(nums []int) int {
	maxVal := max(nums...)
	arr := make([]int, maxVal+1)
	// 下标i位置存放元素i出现的次数之和
	for i := 0; i < len(nums); i++ {
		arr[nums[i]] += nums[i]
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		// 如果要偷当前房间，就不能偷前一间房间
		return max(dfs(i-2)+arr[i], dfs(i-1))
	}
	return dfs(len(arr) - 1)
}

func deleteAndEarnDp(nums []int) int {
	maxVal := max(nums...)
	arr := make([]int, maxVal+1)
	// 下标i位置存放元素i出现的次数之和
	for i := 0; i < len(nums); i++ {
		arr[nums[i]] += nums[i]
	}

	dp := make([]int, len(arr))
	// 处理数组下标越界问题
	// 1.加条件进行判断  2.在dp数组前面多放两个位置下标平移
	dp[0] = arr[0]
	if len(dp) == 1 {
		return dp[0]
	}
	dp[1] = max(arr[0], arr[1])
	if len(dp) == 2 {
		return dp[1]
	}
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-2]+arr[i], dp[i-1])
	}
	return dp[len(dp)-1]
}

func TestDeleteAndEarn(t *testing.T) {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarnDp([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
	fmt.Println(deleteAndEarnDp([]int{2, 2, 3, 3, 3, 4}))
}
