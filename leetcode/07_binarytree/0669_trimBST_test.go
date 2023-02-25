package _7_binarytree

/*
https://leetcode.cn/problems/trim-a-binary-search-tree/description/
*/

/*
模仿450的deleteNode写法需要小心，一定要严格按照后续遍历的方式
*/
func trimBSTError(root *TreeNode, low int, high int) *TreeNode {

	var trimBSTAux func(node *TreeNode, low int, high int) *TreeNode
	trimBSTAux = func(node *TreeNode, low int, high int) *TreeNode {
		if node == nil {
			return nil
		}
		// 如果第一个节点就要删除的话，那么这里的两个递归根本进不去
		if node.Val >= low {
			node.Left = trimBSTAux(node.Left, low, high)
			return node
		} else if node.Val <= high {
			node.Right = trimBSTAux(node.Right, low, high)
			return node
		} else {
			if node.Left == nil && node.Right == nil {
				node = nil
				return nil
			} else if node.Left != nil && node.Right == nil {
				left := node.Left
				node.Left = nil
				return left
			} else if node.Left == nil && node.Right != nil {
				right := node.Right
				node.Right = nil
				return right
			} else {
				left := node.Left
				right := node.Right

				position := node.Right
				for position.Left != nil {
					position = position.Left
				}
				position.Left = left

				node.Left = nil
				node.Right = nil
				return right
			}
		}
	}
	return trimBSTAux(root, low, high)
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {

	var trimBSTAux func(node *TreeNode, low int, high int) *TreeNode
	trimBSTAux = func(node *TreeNode, low int, high int) *TreeNode {
		if node == nil {
			return nil
		}
		// 先处理子树
		node.Left = trimBSTAux(node.Left, low, high)
		node.Right = trimBSTAux(node.Right, low, high)

		if node.Val > high || node.Val < low { // 需要被删除
			if node.Left == nil && node.Right == nil {
				node = nil
				return nil
			} else if node.Left != nil && node.Right == nil {
				left := node.Left
				node.Left = nil
				return left
			} else if node.Left == nil && node.Right != nil {
				right := node.Right
				node.Right = nil
				return right
			} else {
				left := node.Left
				right := node.Right

				position := node.Right
				for position.Left != nil {
					position = position.Left
				}
				position.Left = left

				node.Left = nil
				node.Right = nil
				return right
			}
		} else {
			return node
		}
	}
	return trimBSTAux(root, low, high)
}

// 更加简单的修剪方式
func trimBST2(root *TreeNode, low int, high int) *TreeNode {
	return nil
}
