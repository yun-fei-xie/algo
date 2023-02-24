package _7_binarytree

/*
https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/description/

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
