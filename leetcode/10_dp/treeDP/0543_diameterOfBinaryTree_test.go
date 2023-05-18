package treeDP

import (
	"math"
)

/*
543. 二叉树的直径
https://leetcode.cn/problems/diameter-of-binary-tree/description/?favorite=2cktkvj

给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。

方法：分别计算以当前节点为根节点的左边的最长链和右边的最长链的长度。
在遍历时，不断更新全局的（左边最长链+右边最长链+1）

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree1(root *TreeNode) int {
	max := math.MinInt64

	var findHighest func(node *TreeNode) int
	findHighest = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := findHighest(node.Left)
		right := findHighest(node.Right)

		if left+right > max {
			max = left + right
		}

		if left >= right {
			return left + 1
		} else {
			return right + 1
		}
	}
	findHighest(root)
	return max
}
