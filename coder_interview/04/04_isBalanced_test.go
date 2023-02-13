package _4

/*
*
递归的时候，拿到左右子树的高度，然后进行判断
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, bal := isBalanceAux(root)

	return bal
}

func isBalanceAux(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftTreeHeight, b1 := isBalanceAux(root.Left)
	rightTreeHeight, b2 := isBalanceAux(root.Right)

	height := max(leftTreeHeight, rightTreeHeight) + 1
	if b1 == false || b2 == false {
		return height, false
	}
	if diff(leftTreeHeight, rightTreeHeight) > 1 {
		return height, false
	}
	return height, true

}

func diff(x int, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func max(x, y int) int {

	if x > y {
		return x
	}
	return y
}
