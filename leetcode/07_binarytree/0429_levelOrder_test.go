package _7_binarytree

/*
https://leetcode.cn/problems/n-ary-tree-level-order-traversal/

N叉树的层序遍历
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

解法：相对于二叉树的层序遍历，只是多了一个for循环。



*/
//type Node struct {
//	Val      int
//	Children []*Node
//}
//
//func levelOrderNtree(root *Node) [][]int {
//
//	values := make([][]int, 0)
//
//	var levelorder func(node *Node, depth int)
//	levelorder = func(node *Node, depth int) {
//
//		if node == nil {
//			return
//		}
//		if depth == len(values) {
//			values = append(values, []int{})
//		}
//
//		values[depth] = append(values[depth], node.Val)
//		for i := 0; i < len(node.Children); i++ {
//			child := node.Children[i]
//			levelorder(child, depth+1)
//		}
//
//	}
//
//	levelorder(root, 0)
//	return values
//
//}
