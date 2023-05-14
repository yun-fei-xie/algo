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
513. 找树左下角的值
https://leetcode.cn/problems/find-bottom-left-tree-value/

给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。

方法1：层序遍历-> 我怎么知道自己在最后一层？ 每一层迭代都记住最左边的数值。这个数值一直更新即可
方法2：层序遍历的时候，从右向左进行遍历，这样最后一层的最后一个元素就是要找的答案。
	如何从右向左进行遍历呢？只需要在节点入队的时候，按照从右向左的顺序进行入队。
*/
func findBottomLeftValue1(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	res := 0
	for queue.Len() > 0 {
		l := queue.Len() // 这个地方是一个容易写错的地方
		for i := 0; i < l; i++ {

			node := queue.Remove(queue.Front()).(*TreeNode)
			// 最后一层的第一个节点
			if i == 0 {
				res = node.Val
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return res

}

/*
从右向左进行遍历，不断更新ans
ans的最后一次更新就是答案
*/
func findBottomLeftValue2(root *TreeNode) int {
	ans := -1
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			ans = node.Val
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}
	}
	return ans
}
