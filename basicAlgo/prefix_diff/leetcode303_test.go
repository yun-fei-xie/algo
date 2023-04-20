package prefix_diff

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/range-sum-query-immutable/description/
给定一个整数数组  nums，处理以下类型的多个查询:
计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
实现 NumArray 类：
NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )


*/

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	pre := 0
	prefix := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		prefix[i] = pre + nums[i]
		pre += nums[i]
	}
	return NumArray{prefix: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	//[1 ,2 ,3, 4, 5] -> [1 , 3 , 6 , 10 , 15]
	//
	var rangeLeft, rangeRight int
	if left != 0 {
		rangeLeft = this.prefix[left-1]
	}

	rangeRight = this.prefix[right]

	return rangeRight - rangeLeft
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 *
 */

func TestSumRange(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6}
	numArr := Constructor(arr)
	rangeSum1 := numArr.SumRange(0, 4)
	rangeSum2 := numArr.SumRange(1, 3)
	rangeSum3 := numArr.SumRange(2, 2)
	fmt.Println(rangeSum1)
	fmt.Println(rangeSum2)
	fmt.Println(rangeSum3)
}
