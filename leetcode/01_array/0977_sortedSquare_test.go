package _1_array

import (
	"fmt"
	"sort"
	"testing"
)

/**
https://leetcode.cn/problems/squares-of-a-sorted-array/
*/

/**
1. 乘完直接排序，肯定是最简单，但是性能不太行（排序算法的时间复杂度可以认为是 NlogN）

2. 原数组有序，因此平方后的返回数组的最大值不是在原数组的左边，就是右边。
	可以把两边的数据同时平方并进行比较。将较大值放入结果数组的当前最小值位置。

	定义两个变量i ,j 分别执行nums的首位，定义一个res数组作为结果数组，其长度和nums相同
	定义一个指针r 指向res当前最小值，初始位置为len(res)-1.
	if nums[i]^2 > nums[j]^2 => (i++ , res[r]=nums[i]^2 , r--)
	当i>j时，表示nums数组已经遍历结束。返回结果数组
*/

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

func sortedSquares2(nums []int) []int {
	i := 0
	j := len(nums) - 1
	r := len(nums) - 1
	res := make([]int, len(nums))

	for i <= j {
		if nums[i]*nums[i] >= nums[j]*nums[j] {

			res[r] = nums[i] * nums[i] //先赋值 再修改下标
			i++

		} else {

			res[r] = nums[j] * nums[j]
			j--
		}
		r--
	}
	return res
}

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	res := sortedSquares2(nums)
	fmt.Println(res)
}
