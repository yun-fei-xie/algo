package _7_binarytree

/*
中序+后序 还原二叉树

https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
*/
var hashMap = make(map[int]int, 0)

func buildTree(inorder []int, postorder []int) *TreeNode {
	for index, value := range inorder {
		hashMap[value] = index
	}

	return buildTreeAux(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)

}

func buildTreeAux(inorder []int, postorder []int, inLeft int, inRight int, postLeft int, postRight int) *TreeNode {
	if inLeft > inRight { // 数组区间已经没有元素
		return nil
	}

	root := &TreeNode{Val: postorder[postRight]}
	// 划分区间
	inOrderIndex := hashMap[postorder[postRight]] // 找到后序的root节点在中序数组中的下标

	leftLength := inOrderIndex - inLeft
	rightLength := inRight - inOrderIndex

	root.Left = buildTreeAux(inorder, postorder, inLeft, inOrderIndex-1, postLeft, postLeft+leftLength-1)
	root.Right = buildTreeAux(inorder, postorder, inOrderIndex+1, inRight, postRight-rightLength, postRight-1)

	return root
}
