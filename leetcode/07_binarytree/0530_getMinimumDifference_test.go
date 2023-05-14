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
需要注意的是，求的是最小值，而不是最大值。（如果是最大值的话，用二叉树的最大值-二叉树的最小值）

最小值：两个值最接近的节点。
找两个值最接近的节点-> 排序后，比较相邻节点的差值
排序-> 先序遍历

BST 求最值 一般都是中序遍历
https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
可以放到数组中再求值，也可以直接在遍历中作差
*/
func getMinimumDifference(root *TreeNode) int {
	var result = math.MaxInt64
	var pre *TreeNode = nil
	var inOrder func(node *TreeNode)

	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)
		// 对于第一个节点，进行特殊判断
		if pre != nil {
			result = getMin(result, pre.Val, node.Val)
		}

		pre = node

		inOrder(node.Right)
	}

	inOrder(root)
	return result
}

/*
取 r1 与 abs（r2 , r3）的较小值
*/
func getMin(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}
