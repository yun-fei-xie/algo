package _7_binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
1026. 节点与其祖先之间的最大差值
https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/description/

给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。
（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
节点数量最小为2个。

方法：使用前序遍历，记录从根节点到当前节点[root...currentNode]路径上的最大值和最小值
用最大值-最小值就可以得到一个diff。把所有的diff进行比较，就可以得到全局的最大差值。
*/
func maxAncestorDiff(root *TreeNode) int {

	var maxDiff = 0
	var preOrder func(node *TreeNode, maxVal int, minVal int)
	preOrder = func(node *TreeNode, maxVal int, minVal int) {
		// node不会为空，因为在递归过程中控制不让进入空节点。
		if node.Val > maxVal {
			maxVal = node.Val
		}
		if node.Val < minVal {
			minVal = node.Val
		}

		maxDiff = max(maxDiff, maxVal-minVal)

		if node.Left != nil {
			preOrder(node.Left, maxVal, minVal)
		}
		if node.Right != nil {
			preOrder(node.Right, maxVal, minVal)
		}

	}

	if root.Left != nil {
		preOrder(root.Left, root.Val, root.Val)
	}

	if root.Right != nil {
		preOrder(root.Right, root.Val, root.Val)
	}
	return maxDiff
}
