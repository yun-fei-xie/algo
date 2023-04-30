package jianzhi_offer

import (
	"fmt"
	"testing"
)

/*
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
示例 1：
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

nums -> [2, 3, 1, 0, 2, 5, 3]
sort -> [0, 1, 2, 2, 3, 3, 5]

nums -> [3, 1, 2, 3]

sort -> [1, 2, 3 ,3]

最简单的方式是遍历一遍数组，用hashSet在遍历的时候查找是否当前元素在前面出现过。
时间复杂度O(n),空间复杂度O(n)
*/
func findRepeatNumber(nums []int) int {
	set := make(map[int]struct{})
	length := len(nums)
	for i := 0; i < length; i++ {
		if _, found := set[nums[i]]; found {
			return nums[i]
		} else {
			set[nums[i]] = struct{}{}
		}
	}
	return -1
}

func TestFindRepeatNumber(t *testing.T) {
	fmt.Println(findRepeatNumber([]int{3, 1, 2, 3}))
}
