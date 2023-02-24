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

	root.Left = constructMaximumAux(nums[:maxIndex])
	root.Right = constructMaximumAux(nums[maxIndex+1:])

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
