package _4_design

import (
	"fmt"
	"math/rand"
	"testing"
)

/*
https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn6gq1/
给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。

实现 Solution class:

Solution(int[] nums) 使用整数数组 nums 初始化对象
int[] reset() 重设数组到它的初始状态并返回
int[] shuffle() 返回数组随机打乱后的结果
*/
type Solution struct {
	arr []int
}

func Constructor(nums []int) Solution {
	return Solution{arr: nums}
}

func (this *Solution) Reset() []int {
	return this.arr
}

/*
Knuth-Shuffle 算法
*/
func (this *Solution) Shuffle() []int {
	arrCopy := make([]int, len(this.arr))
	copy(arrCopy, this.arr)
	for i := len(arrCopy) - 1; i >= 0; i-- {
		randIndex := rand.Intn(i + 1) //[0 ,i+1)
		arrCopy[i], arrCopy[randIndex] = arrCopy[randIndex], arrCopy[i]
	}
	return arrCopy
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func TestSolution(t *testing.T) {
	fmt.Println()
}
