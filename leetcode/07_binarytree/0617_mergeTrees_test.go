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
https://leetcode.cn/problems/merge-two-binary-trees/description/
*/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return mergeTreesAux(root1, root2)
}

// 将root2合并到root1
func mergeTreesAux(node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if node1 == nil {
		return node2
	} else if node2 == nil {
		return node1
	}

	node1.Val += node2.Val
	node1.Left = mergeTreesAux(node1.Left, node2.Left)
	node1.Right = mergeTreesAux(node1.Right, node2.Right)

	return node1
}
