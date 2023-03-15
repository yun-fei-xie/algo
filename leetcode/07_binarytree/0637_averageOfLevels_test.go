package _7_binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
https://leetcode.cn/problems/average-of-levels-in-binary-tree/description/

给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。

解法：层序遍历将同一层的元素放入到同一个数组中，最后对每一个数组求解平均值。

*/
func averageOfLevels(root *TreeNode) []float64 {

	values := make([][]int, 0)
	res := make([]float64, 0)

	var levelOrder func(node *TreeNode, depth int)
	levelOrder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(values) == depth {
			values = append(values, []int{})
		}

		values[depth] = append(values[depth], node.Val)
		levelOrder(node.Left, depth+1)
		levelOrder(node.Right, depth+1)
	}
	levelOrder(root, 0)

	for i := 0; i < len(values); i++ {
		res = append(res, avg(values[i]))
	}
	return res
}

func avg(arr []int) float64 {
	var sum = 0
	for _, num := range arr {
		sum += num
	}
	return float64(sum) / float64(len(arr))
}
