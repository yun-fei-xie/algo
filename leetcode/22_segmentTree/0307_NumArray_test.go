package _2_segmentTree_test

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

方法：树状数组（树状数组求解区间查询问题转化为求解两次前缀和问题）
方法：线段树 这个文件中使用线段树实现
*/

type NumArray struct {
	arr  []int
	tree []int
}

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums))
	copy(arr, nums)
	tree := make([]int, 4*len(nums))

	buildSegmentTree(tree, arr, 0, 0, len(nums)-1)
	return NumArray{
		arr:  arr,
		tree: tree,
	}
}

func buildSegmentTree(tree []int, data []int, treeIndex int, left int, right int) {
	if left == right {
		tree[treeIndex] = data[left]
		return
	}
	mid := left + (right-left)/2
	leftChild := leftChildIndex(treeIndex)
	rightChild := rightChildIndex(treeIndex)

	buildSegmentTree(tree, data, leftChild, left, mid)
	buildSegmentTree(tree, data, rightChild, mid+1, right)
	// 区间求和
	tree[treeIndex] = tree[leftChild] + tree[rightChild]
}

func leftChildIndex(idx int) int {
	return 2*idx + 1
}
func rightChildIndex(idx int) int {
	return 2*idx + 2
}

func (this *NumArray) Update(index int, val int) {
	this.arr[index] = val
	this.update(0, 0, len(this.arr)-1, index, val)
}

func (this *NumArray) update(treeIndex int, left int, right int, index int, val int) {
	if left == right && left == index {
		this.tree[treeIndex] = val
		return
	}
	mid := left + (right-left)/2
	leftChild := leftChildIndex(treeIndex)
	rightChild := rightChildIndex(treeIndex)

	if index <= mid {
		this.update(leftChild, left, mid, index, val)
		this.tree[treeIndex] = this.tree[leftChild] + this.tree[rightChild]
	}
	if index > mid {
		this.update(rightChild, mid+1, right, index, val)
		this.tree[treeIndex] = this.tree[leftChild] + this.tree[rightChild]
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sumRange(0, 0, len(this.arr)-1, left, right)
}
func (this *NumArray) sumRange(treeIndex int, left int, right int, queryLeft int, queryRight int) int {
	if queryLeft == left && queryRight == right {
		return this.tree[treeIndex] // 注意这里修改的是treeIndex位置的值
	}
	mid := left + (right-left)/2
	leftChild := leftChildIndex(treeIndex)
	rightChild := rightChildIndex(treeIndex)

	if queryLeft > mid {
		return this.sumRange(rightChild, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		return this.sumRange(leftChild, left, mid, queryLeft, queryRight)
	}

	return this.sumRange(leftChild, left, mid, queryLeft, mid) + this.sumRange(rightChild, mid+1, right, mid+1, queryRight)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
func TestNumArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	nm := Constructor(arr)
	fmt.Println(nm.SumRange(0, 7))
	fmt.Println(nm.SumRange(4, 7))
	nm.Update(0, 100)
	fmt.Println(nm.SumRange(0, 7))
	fmt.Println(nm.SumRange(4, 7))
}
