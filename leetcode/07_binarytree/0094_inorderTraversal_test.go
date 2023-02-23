package _7_binarytree

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
二叉树中序遍历 同样使用内部函数
*/
func inorderTraversal(root *TreeNode) []int {

	values := make([]int, 0)

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {

		if node == nil {
			return
		}

		inorder(node.Left)
		values = append(values, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return values

}

// 迭代法 中序遍历

/*
中序遍历的顺序是 左-中-右。当前节点的左子树必须全部处理完毕才能处理当前节点。
于是需要一种方式知道当前节点的左子树是不是都处理完了
用一个指针，从根节点 沿着左子树 一路向下
*/

func inorderTraversalIter(root *TreeNode) []int {
	values := make([]int, 0)
	stack := list.New()
	if root == nil {
		return values
	}

	cur := root

	for stack.Len() > 0 || cur != nil {

		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			node := stack.Remove(stack.Back()).(*TreeNode)
			values = append(values, node.Val)
			cur = node.Right
		}

	}
	return values
}
