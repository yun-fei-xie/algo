package treeDP

import (
	"fmt"
	"testing"
)

/*
95. 不同的二叉搜索树 II
https://leetcode.cn/problems/unique-binary-search-trees-ii/description/
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
方法：在[left,right]这个区间中，枚举以i为根节点的子树。并将结果集合返回给上层
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	// 返回以[left,right]这段数字构建的BST的树的根节点集合
	var dfs func(left int, right int) []*TreeNode
	dfs = func(left int, right int) []*TreeNode {
		if left > right {
			// return nil 不行因为上层拿不到左右子树不会进入循环 于是上上层也不会拿到子树节点
			return []*TreeNode{nil}
		}
		var ans = make([]*TreeNode, 0)
		for i := left; i <= right; i++ {
			//以i为根节点
			leftChildren := dfs(left, i-1)
			rightChildren := dfs(i+1, right)

			for _, leftChild := range leftChildren {
				for _, rightChild := range rightChildren {
					rootNode := genNode(i)
					rootNode.Left = leftChild
					rootNode.Right = rightChild
					ans = append(ans, rootNode)
				}
			}
		}
		return ans
	}
	return dfs(1, n)
}

func genNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func TestGenerateTrees(t *testing.T) {
	fmt.Println(generateTrees(3))
}
