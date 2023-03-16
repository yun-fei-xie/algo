package _7_binarytree

/*
https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/description/

给定一个 N 叉树，找到其最大深度。
最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。

*/

type Node struct {
	Val      int
	Children []*Node
}

func maxDepthN(root *Node) int {
	return maxDepthNAux(root)
}

func maxDepthNAux(node *Node) int {
	if node == nil {
		return 0
	}

	max := 0 // 高度最小也就是0 （之前用 int.Min 造成了错误）
	for i := 0; i < len(node.Children); i++ {
		height := maxDepthNAux(node.Children[i])
		if height > max {
			max = height
		}
	}
	return max + 1
}
