package _4

/**
给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

示例:
给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0
         / \
       -3   9
       /   /
     -10  5
*/

/**
思路：感觉从中间节点作为树的根节点，
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else {
		return arrayToBSTAux(nums, 0, len(nums)-1)
	}
}

func arrayToBSTAux(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil

	} else {
		index := (left + right) / 2
		rootVal := nums[index]
		rootNode := &TreeNode{
			Val:   rootVal,
			Left:  nil,
			Right: nil,
		}

		rootNode.Left = arrayToBSTAux(nums, left, index-1)
		rootNode.Right = arrayToBSTAux(nums, index+1, right)
		return rootNode
	}
}
