package _7_binarytree

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
递归+回溯的一个题 仔细体会
*/
func binaryTreePaths(root *TreeNode) []string {

	res := make([]string, 0)

	var preorder func(node *TreeNode, path string)
	preorder = func(node *TreeNode, path string) {

		if node.Left == nil && node.Right == nil {
			path = path + strconv.Itoa(node.Val)
			res = append(res, path)
		}

		path = path + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			preorder(node.Left, path)
		}

		if node.Right != nil {
			preorder(node.Right, path)
		}

	}
	preorder(root, "")
	return res
}
