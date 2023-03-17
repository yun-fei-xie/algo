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

https://leetcode.cn/problems/binary-tree-paths/

二叉树的所有路径
从根节点到所有叶子节点的路径

第一反应其实是用一个数组来存储单个路径，但是那样未免过于麻烦
直接用一个字符串更加方便省事儿

*/
func binaryTreePaths(root *TreeNode) []string {

	res := make([]string, 0)

	var preorder func(node *TreeNode, path string)
	preorder = func(node *TreeNode, path string) {

		if node.Left == nil && node.Right == nil { // 递归到底,拿到了一条路径
			path = path + strconv.Itoa(node.Val)
			res = append(res, path)
			return
		}

		path = path + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			preorder(node.Left, path)
		}

		if node.Right != nil {
			preorder(node.Right, path)
		}

	}
	preorder(root, "") // 这里的path是字符串，值传递 所以不需要做字符回退处理。
	return res
}
