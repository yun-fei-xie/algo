package _3_hashtable

/*
https://leetcode.cn/problems/two-sum/description/
**数组中同一个元素在答案里不能重复出现** 表示同一个数字不能使用2次

*1. 二重循环
 2. hash 表
 3. 如果是有序数组 是否可以使用滑动窗口
*/
func twoSum(nums []int, target int) []int {

	res := make([]int, 0)
loop:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				res = append(res, i)
				res = append(res, j)
				break loop
			}
		}
	}
	return res
}

/*

使用map 查找target对当前元素的差值
将原本的时间复杂度从 n^2 降低到了 n * logn

*/
func twoSum2(nums []int, target int) []int {
	numsMap := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}

loop:
	for index := 0; index < len(nums); index++ {
		if index2, found := numsMap[target-nums[index]]; found && index != index2 {
			res = append(res, index)
			res = append(res, index2)
			break loop
		}
	}
	return res

}
