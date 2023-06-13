package _3_QuadTree

import "testing"

/*

558. 四叉树交集
https://leetcode.cn/problems/logical-or-of-two-binary-grids-represented-as-quad-trees/description/


方法：递归求解
1. 如果两个同级节点同时都是叶子节点。那么交集就是一个叶子节点。叶子节点的值是两个叶子节点代表的值的或|运算的结果。只要有一个是1生成的节点就是1。
2. 如果两个同级节点一个是叶子节点，一个不是叶子节点。如果叶子节点代表0区域，那么结果是叶子节点0区域；如果叶子节点是1区域，那么结果是非叶子节点。
3. 如果两个节点都不是叶子节点，继续递归。
*/

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

/*
这个版本的解答有问题，回头重写
*/
func intersect(quadTree1 *Node, quadTree2 *Node) *Node {

	var helper func(root1 *Node, root2 *Node) *Node
	helper = func(root1 *Node, root2 *Node) *Node {
		// 两个节点同时是叶子节点
		if root1.IsLeaf && root2.IsLeaf {
			node := &Node{
				Val:    false,
				IsLeaf: true,
			}
			if root1.Val || root2.Val {
				node.Val = true
			}
			return node
		}
		// 一个是叶子节点 一个不是叶子节点
		if root1.IsLeaf || root2.IsLeaf {
			// 有一个是叶子节点 并且叶子节点代表0
			if root1.IsLeaf && root1.Val == false || root2.IsLeaf && root2.Val == false {
				return &Node{
					Val:    false,
					IsLeaf: true,
				}
			}
			// 有一个是叶子节点，并且叶子节点代表1
			if root1.IsLeaf && root1.Val == true {
				// 怎么深度copy
				return root2

			} else if root2.IsLeaf && root2.Val == true {
				return root1
			}
		}

		// 两个都不是叶子节点(注意：两个非叶子节点可能会产生叶子节点（0 1互补))

		topLeft := helper(root1.TopLeft, root2.TopLeft)
		topRight := helper(root1.TopRight, root2.TopRight)
		bottomLeft := helper(root1.BottomLeft, root2.BottomLeft)
		bottomRight := helper(root1.BottomRight, root2.BottomRight)

		if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf && topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
			return &Node{
				Val:         topRight.Val,
				IsLeaf:      true,
				TopLeft:     topLeft,
				TopRight:    topRight,
				BottomLeft:  bottomLeft,
				BottomRight: bottomRight,
			}
		}

		return &Node{
			Val:         false,
			IsLeaf:      false,
			TopLeft:     topLeft,
			TopRight:    topRight,
			BottomLeft:  bottomLeft,
			BottomRight: bottomRight,
		}
	}
	return helper(quadTree1, quadTree2)
}

func TestIntersect(t *testing.T) {
}
