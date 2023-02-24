package _7_binarytree

/*
https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/

填充每个节点的下一个右侧节点指针
*/

//type Node struct {
//	Val   int
//	Left  *Node
//	Right *Node
//	Next  *Node
//}
//
//func connect(root *Node) *Node {
//	nodes := make([][]*Node, 0)
//
//	var levelOrder func(node *Node, depth int)
//	levelOrder = func(node *Node, depth int) {
//		if node == nil {
//			return
//		}
//
//		if len(nodes) == depth {
//			nodes = append(nodes, []*Node{})
//		}
//		nodes[depth] = append(nodes[depth], node)
//		levelOrder(node.Left, depth+1)
//		levelOrder(node.Right, depth+1)
//
//	}
//	levelOrder(root, 0)
//
//	// 最后把每一层的节点取出 进行串联
//	for i := 0; i < len(nodes); i++ {
//		for j := 0; j < len(nodes[i])-1; j++ {
//			nodes[i][j].Next = nodes[i][j+1]
//		}
//	}
//	return root
//}
