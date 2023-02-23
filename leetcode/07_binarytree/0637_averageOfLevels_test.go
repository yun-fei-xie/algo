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
