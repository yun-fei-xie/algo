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
https://leetcode.cn/problems/search-in-a-binary-search-tree/description/

在二叉搜索树中搜索一个元素

*/
func searchBST(root *TreeNode, val int) *TreeNode {
	return searchBSTAux(root, val)
}

func searchBSTAux(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return node // return nil   node==nil || node.Val==Val 这两个条件可以合并
	}

	if node.Val == val {
		return node
	}

	if node.Val > val {
		return searchBST(node.Left, val)
	} else {
		return searchBST(node.Right, val)
	}
}
