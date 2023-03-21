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
https://leetcode.cn/problems/merge-two-binary-trees/description/

给你两棵二叉树： root1 和 root2 。
想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
返回合并后的二叉树。
注意: 合并过程必须从两个树的根节点开始。


解题思路：

这个题背后的考察点应该是：如何同时遍历两棵树？ 和这道题很相似的一题镜像二叉树的判断。


这里假设把root2 合并到root1上



*/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return mergeTreesAux(root1, root2)
}

// 将root2合并到root1
func mergeTreesAux(node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if node1 == nil {
		return node2
	} else if node2 == nil {
		return node1
	}

	node1.Val += node2.Val
	node1.Left = mergeTreesAux(node1.Left, node2.Left)
	node1.Right = mergeTreesAux(node1.Right, node2.Right)

	return node1
}
