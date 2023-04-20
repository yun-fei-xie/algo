package _7_binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/
*/

//var m = make(map[int]int)

func sortedArrayToBST(nums []int) *TreeNode {

	//for index, num := range nums {
	//	m[num] = index
	//}

	var arrayToBstAux func(nums []int, left, right int) *TreeNode

	arrayToBstAux = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		rootIndex := (left + right) / 2
		root := &TreeNode{Val: nums[rootIndex]}
		root.Left = arrayToBstAux(nums, left, rootIndex-1)
		root.Right = arrayToBstAux(nums, rootIndex+1, right)
		return root
	}

	return arrayToBstAux(nums, 0, len(nums)-1)
}
