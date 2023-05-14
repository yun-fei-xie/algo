package _7_binarytree

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/
129. 求根节点到叶节点数字之和

给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：
例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。
叶节点 是指没有子节点的节点。

方法：字符串求和+前序遍历
注意点：在叶子节点的时候收集答案，不要去空节点。

*/
func sumNumbers(root *TreeNode) int {
	var ans int
	var preOrder func(node *TreeNode, number int)
	preOrder = func(node *TreeNode, number int) {
		if node != nil {
			number = number*10 + node.Val
		}
		if node.Left == nil && node.Right == nil {
			ans += number
			return
		}

		if node.Left != nil {
			preOrder(node.Left, number)
		}
		if node.Right != nil {
			preOrder(node.Right, number)
		}
	}
	preOrder(root, 0)
	return ans

}

func TestSumNumbers(t *testing.T) {

	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3}}

	fmt.Println(sumNumbers(root))
}
