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
https://leetcode.cn/problems/delete-node-in-a-bst/description/
1. 找到要删除的节点
2. 删除节点
	1. 叶子节点
	2. 非叶子节点
		 左孩子为空
         右孩子为空
         左右都不为空
	当左右都不为空的时候，需要找到右子树中最左边的节点，然后把左子树当做其左孩子挂载上去。
*/
func deleteNode(root *TreeNode, key int) *TreeNode {

	var deleteAux func(node *TreeNode, key int) *TreeNode
	deleteAux = func(node *TreeNode, key int) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val > key { //left
			node.Left = deleteAux(node.Left, key)
			return node
		} else if node.Val < key {
			node.Right = deleteAux(node.Right, key)
			return node
		} else { // 当前节点
			if node.Left == nil && node.Right == nil { // 叶子节点
				return nil
			} else if node.Left != nil && node.Right == nil { // 左孩子存在
				left := node.Left
				node.Left = nil
				return left
			} else if node.Right != nil && node.Left == nil {
				right := node.Right
				node.Right = nil
				return right
			} else { // 左右都不为空 右孩子向上提、左孩子挂上去（找到右子树最小的那个元素，把左子树挂上去）
				left := node.Left
				right := node.Right

				node.Left = nil
				position := node.Right
				for position.Left != nil {
					position = position.Left
				}
				position.Left = left
				node.Right = nil
				return right
			}
		}

	}
	return deleteAux(root, key)
}
