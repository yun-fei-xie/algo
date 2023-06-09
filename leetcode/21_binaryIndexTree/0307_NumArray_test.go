package _1_binaryIndexTree

import (
	"fmt"
	"testing"
)

/*
307. 区域和检索 - 数组可修改
https://leetcode.cn/problems/range-sum-query-mutable/description/

给你一个数组 nums ，请你完成两类查询。
其中一类查询要求 更新 数组 nums 下标对应的值
另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
实现 NumArray 类：
NumArray(int[] nums) 用整数数组 nums 初始化对象
void update(int index, int val) 将 nums[index] 的值 更新 为 val
int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）

方法：树状数组
*/
type NumArray struct {
	bit []int
	num []int
}

func Constructor(nums []int) NumArray {
	length := len(nums)
	bit := make([]int, length+1)
	for i := 0; i < length; i++ {
		update(bit, i+1, nums[i])
	}
	return NumArray{bit: bit, num: nums}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.num[index]
	this.num[index] = val
	update(this.bit, index+1, diff)
}
func update(bit []int, index int, val int) {
	for index < len(bit) {
		bit[index] += val
		index += index & (-index)
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.getSum(right+1) - this.getSum(left)
}

func (this *NumArray) getSum(index int) (ans int) {
	for index != 0 {
		ans += this.bit[index]
		index -= index & (-index)
	}
	return ans
}

func TestBIT(t *testing.T) {
	bit := Constructor([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}) //10
	fmt.Println(bit.SumRange(0, 9))
	fmt.Println(bit.SumRange(4, 6))
	bit.Update(5, 10)
	bit.Update(3, 7)
	fmt.Println(bit.SumRange(0, 8))

}
