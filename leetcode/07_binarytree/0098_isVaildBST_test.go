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

98. 验证二叉搜索树

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

对这一题理解有误，不能仅仅比较当前节点的大于左孩子 小于右孩子 这是不对的。
为什么？ 下面👇🏻这棵二叉树不满足定义。

		5
      /   \
	4      6
         /   \
		3    7


binaryTree 中序遍历应该有序（后面一个被遍历到的值 应该大于前面的）

一方面可以使用数组收集中序遍历的值 ，然后判断数组是否有序
另一方面也可以在遍历中直接判断。
如果发现不符合条件的节点，直接将全局flag设置为false

方法2：前序遍历，在递归的过程中，固定孩子节点值的左右边界。
当递归到下一层时，检查当前节点的值是否在左右边界中。



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
区间范围方法
每一个节点都有一个数值范围，每一个节点的数值范围会受到父节点的限制。
*/

func isValidBST2(root *TreeNode) bool {

	var preOrder func(node *TreeNode, leftBound int, rightBound int) bool
	preOrder = func(node *TreeNode, leftBound int, rightBound int) bool {
		if node == nil {
			return true
		}
		if node.Val <= leftBound || node.Val >= rightBound {
			return false
		}
		return preOrder(node.Left, leftBound, node.Val) && preOrder(node.Right, node.Val, rightBound)
	}

	return preOrder(root, math.MinInt, math.MaxInt)
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
