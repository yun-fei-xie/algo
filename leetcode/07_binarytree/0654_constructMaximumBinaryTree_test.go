package _7_binarytree

import (
	"fmt"
	"math"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
https://leetcode.cn/problems/maximum-binary-tree/description/

给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:
创建一个根节点，其值为 nums 中的最大值。
递归地在最大值 左边 的 子数组前缀上 构建左子树。
递归地在最大值 右边 的 子数组后缀上 构建右子树。
返回 nums 构建的 最大二叉树 。

这个题比较有意思的一点：右子树的值不一定会全部大于左子树的值

每次从arr[left , right]找出最大的一个值的下标 maxIndex ,
然后将这个区间一份为二：left-> [left, maxIndex-1]   right -> [maxIndex+1 , right]

*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumAux(nums)
}

func constructMaximumAux(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex, max := findMax(nums)
	root := &TreeNode{
		Val: max,
	}

	root.Left = constructMaximumAux(nums[:maxIndex])    // 递归从左边进行查找
	root.Right = constructMaximumAux(nums[maxIndex+1:]) // 递归从右边进行查找

	return root
}

func findMax(arr []int) (index int, num int) {
	max := math.MinInt64
	maxIndex := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			maxIndex = i
			max = arr[i]
		}
	}
	return maxIndex, max
}

func TestIndex(t *testing.T) {

	arr := []int{1, 2, 3}
	fmt.Println(arr[:2]) // [ ... ) 左闭右开

}
