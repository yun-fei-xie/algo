package easy

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/description/?favorite=2cktkvj

给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

思路：鸽笼原理 , n个位置，对nums数组进行一次比那里后可以知道哪些位置没有被占用。
	第二次对position进行遍历，收集没有被占用的位置的下标。

*/

func findDisappearedNumbers(nums []int) []int {
	position := make([]bool, len(nums)+1)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		position[nums[i]] = true
	}

	for i := 1; i < len(position); i++ {
		if position[i] == false {
			res = append(res, i)
		}
	}
	return res
}

func TestFindDisappered(t *testing.T) {
	res := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(res)
}
