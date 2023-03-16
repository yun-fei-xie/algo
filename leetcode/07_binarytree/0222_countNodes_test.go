package _7_binarytree

import "container/list"

/*
https://leetcode.cn/problems/count-complete-tree-nodes/description/

给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

如何不使用遍历，把时间复杂度压缩到O(1)
*/

/*
广度优先遍历解法
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	totalCounts := 0
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {

		node := queue.Remove(queue.Front()).(*TreeNode)
		totalCounts++
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return totalCounts
}

/*
递归解法,极简
但是没有用到完全二叉树的性质
*/

func countNodes2(root *TreeNode) int {

	var countNodes func(node *TreeNode) int
	countNodes = func(node *TreeNode) int {
		if node == nil {
			return 0
		} else {
			return countNodes(node.Left) + countNodes(node.Right) + 1

		}
	}
	return countNodes(root)
}

/*
使用完全二叉树的性质

规定根节点位于第0层，完全二叉树的最大层数为h。
根据完全二叉树的特性可知，完全二叉树的最左边的节点一定位于最底层。
因此从根节点出发，每次访问左子节点，直到遇到叶子节点，该叶子节点即为完全二叉树的最左边的节点，经过的路径长度即为最大层数h。
当h0≤i<h 时，第i层包含2^i个节点，最底层包含的节点数最少为1，最多为2^h

题解参考:https://leetcode.cn/problems/count-complete-tree-nodes/solutions/495655/wan-quan-er-cha-shu-de-jie-dian-ge-shu-by-leetco-2/
*/
func countNode3(root *TreeNode) int {
	//todo
	return 0
}
