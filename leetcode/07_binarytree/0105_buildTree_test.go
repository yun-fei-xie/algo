package _7_binarytree

/*
https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
最重要的就是确认下标
如果思路没问题的情况下出错，使用小数据量检查下标
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
