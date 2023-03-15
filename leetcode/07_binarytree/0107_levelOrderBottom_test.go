package _7_binarytree

/*
https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/



给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

递归层序遍历
最后翻转数组
*/

func levelOrderBottom(root *TreeNode) [][]int {

	values := [][]int{}

	var levelorder func(node *TreeNode, depth int)

	levelorder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		// 问题出在这里，如果能够在返回的时候创建数组，就可以实现内层数组从bottom->up的顺序
		if len(values) == depth {
			values = append(values, []int{})
		}

		values[depth] = append(values[depth], node.Val)

		levelorder(node.Left, depth+1)
		levelorder(node.Right, depth+1)

	}

	levelorder(root, 0)

	// 把数组翻转一下
	i := 0
	j := len(values) - 1
	for i < j {
		values[i], values[j] = values[j], values[i]
		i++
		j--
	}

	return values

}
