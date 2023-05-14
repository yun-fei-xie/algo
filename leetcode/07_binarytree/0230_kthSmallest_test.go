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
https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？

方法1：中序遍历有序，使用中序遍历拿到第k小的值
使用递归的方法会遍历整棵树，有额外的不要的遍历开销。
因此，可以使用迭代的方法。

*/

// 递归
func kthSmallest(root *TreeNode, k int) int {
	var ans int
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		k--
		if k == 0 {
			ans = node.Val
		}
		inOrder(node.Right)
	}

	inOrder(root)
	return ans
}

//迭代

func kthSmallest2(root *TreeNode, k int) int {

	stack := list.New()
	cur := root

	for stack.Len() > 0 || cur != nil {

		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			node := stack.Remove(stack.Back()).(*TreeNode)
			k--
			if k == 0 {
				return node.Val
			}
			cur = node.Right
		}

	}
	return -1
}
