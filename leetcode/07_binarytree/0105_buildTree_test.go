package _7_binarytree

/*
https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

从前序和中序遍历构造二叉树

解法：前序可以确定根节点，然后用前序中的根节点的位置在中序数组中进行左右子树的划分。
在前序中找到根节点的值后，需要在中序数组中找到这个值的index。为了不需要每次都进行遍历，提前用一个map把inorder中的索引保存起来。

*/

var hashM = make(map[int]int)

func buildTreePI(preorder []int, inorder []int) *TreeNode {
	for index, num := range inorder {
		hashM[num] = index
	}

	return buildTreeFromPIAux(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTreeFromPIAux(preorder []int, inorder []int, preLeft int, preRight int, inLeft int, inRight int) *TreeNode {
	if preLeft > preRight {
		return nil
	}

	root := &TreeNode{Val: preorder[preLeft]}
	rootIndexInOrder := hashM[preorder[preLeft]]
	leftLength := rootIndexInOrder - inLeft
	rightLength := inRight - rootIndexInOrder

	root.Left = buildTreeFromPIAux(preorder, inorder, preLeft+1, preLeft+leftLength, inLeft, rootIndexInOrder-1)
	root.Right = buildTreeFromPIAux(preorder, inorder, preRight-rightLength+1, preRight, rootIndexInOrder+1, inRight)

	return root
}
