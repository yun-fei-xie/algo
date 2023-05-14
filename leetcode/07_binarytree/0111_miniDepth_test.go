package _7_binarytree_test

import "container/list"

/*
https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

解法：
1. 首先可以想到使用深度优先搜索的方法，遍历整棵树，记录最小深度。
对于每一个非叶子节点，我们只需要分别计算其左右子树的最小叶子节点深度。这样就将一个大问题转化为了小问题，可以递归地解决该问题。
递归写法需要考虑链表那种情况，不然很容易写出错误的代码。
例如：如果左孩子为空，那么应该去看一下右孩子，如果右孩子为空，应该去看一下左孩子，而不是直接返回

		3
	  /    \

9        20

	           \
		        7

例如，在20这个子树，会返回高度为2，但是在3这个位置会返回9的子树高度+1

2.层序遍历
在层序遍历的时候，由于是自顶向下进行遍历。
因此，第一个被发现的叶子节点所在的深度就是最小深度。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepthRec(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var min func(i, j int) int
	min = func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 1
		}
		if node.Left == nil && node.Right != nil {
			return dfs(node.Right) + 1
		}
		if node.Right == nil && node.Left != nil {
			return dfs(node.Left) + 1
		}

		return min(dfs(node.Left), dfs(node.Right)) + 1
	}
	return dfs(root)
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	minDep := 0

	for queue.Len() > 0 {

		levelCount := queue.Len()
		minDep++
		for i := 0; i < levelCount; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return minDep
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}
	}
	return minDep
}
