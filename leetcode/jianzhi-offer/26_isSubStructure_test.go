package jianzhi_offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	var preTravels func(nodeA *TreeNode, nodeB *TreeNode) bool
	preTravels = func(nodeA *TreeNode, nodeB *TreeNode) bool {
		if nodeA == nil {
			return false
		}
		if contains(nodeA, nodeB) {
			return true
		}
		return preTravels(nodeA.Left, nodeB) || preTravels(nodeA.Right, nodeB)
	}
	return preTravels(A, B)
}

/*
contains 以nodeA为根节点的树是否包含以nodeB为根节点的树
*/
func contains(nodeA *TreeNode, nodeB *TreeNode) bool {
	if nodeB == nil {
		return true
	}
	if nodeA == nil || nodeA.Val != nodeB.Val {
		return false
	}
	return contains(nodeA.Left, nodeB.Left) && contains(nodeA.Right, nodeB.Right)
}
