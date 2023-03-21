package _7_binarytree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
对这一题理解有误，不能仅仅比较当前节点的大于左孩子 小于右孩子 这是不对的。
为什么？ 下面👇🏻这棵二叉树不满足定义。

		5
      /   \
	4      6
         /   \
		3    7


bst 中序遍历应该有序（后面一个被遍历到的值 应该大于前面的）

一方面可以使用数组收集中序遍历的值 ，然后判断数组是否有序
另一方面也可以在遍历中直接判断。
如果发现不符合条件的节点，直接将全局flag设置为false

*/
func isValidBST(root *TreeNode) bool {
	minInt := math.MinInt64
	res := true

	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)
		if node.Val > minInt {
			minInt = node.Val
		} else {
			res = false
		}

		inOrder(node.Right)
	}
	inOrder(root)
	return res

}

/*
错误的解答
*/
func isValidBSTAux(node *TreeNode) bool {
	if node == nil {
		return true
	}
	// 叶子
	if node.Left == nil && node.Right == nil {
		return true
	}

	left := isValidBSTAux(node.Left)
	right := isValidBSTAux(node.Right)

	var leftValid = true
	var rightValid = true
	if node.Left != nil && node.Left.Val >= node.Val {
		leftValid = false
	}
	if node.Right != nil && node.Right.Val <= node.Val {
		rightValid = false
	}

	return left && right && leftValid && rightValid

}
