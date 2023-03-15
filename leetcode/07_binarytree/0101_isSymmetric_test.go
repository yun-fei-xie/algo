package _7_binarytree

/*
https://leetcode.cn/problems/symmetric-tree/

对比一颗树的左右子树是否对称

                         1
                     2         2
                  3    4     4    3
                5  7  8 10 10 8  7 5

这个题的难点在于，使用递归的方法如何比较不在同一个分支上的节点呢？
比如上图中，最后一行的两个节点10。他们所在的子树都不一样啊。
例如：如何比较最底层的对称的两个节点10呢？
做法：当比较第二层的两个节点2时，向下递归比较左边2的右节点（左边的节点4） 和右边2的左节点（右边的节点4）
于是当前递归函数就来到了2个节点4。通过这样的方式，便可以比较不在同一个子树上的节点。


*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compareTree(root.Left, root.Right)

}

func compareTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}

	// left.val == right.val

	boolLeft := compareTree(left.Left, right.Right)
	boolRight := compareTree(left.Right, right.Left)

	return boolLeft && boolRight
}
