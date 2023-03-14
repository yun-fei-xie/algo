package mid

import "container/list"

/*
https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/?favorite=2cktkvj

给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。


思路：
第一种方法：用一个list收集前序遍历过程中遇到的节点。前序遍历结束后，从list中从前到后取出节点再挨个串起来。此方法需要O(n)的方法

第二种方法：通过后序遍历



*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten1(root *TreeNode) {
	if root == nil {
		return
	}

	var preOrderTrave func(node *TreeNode)
	l := list.New()
	preOrderTrave = func(node *TreeNode) {
		if node == nil {
			return
		}

		l.PushBack(node)
		preOrderTrave(node.Left)
		preOrderTrave(node.Right)
	}
	preOrderTrave(root)
	cur := l.Front().Value.(*TreeNode) // 第一个节点
	l.Remove(l.Front())

	// 这个地方处理的时候需要📢注意一点：next.Left=nil next.Right=nil 用来保证末尾节点可以指向nil

	for l.Len() != 0 {
		node := l.Front()
		l.Remove(node)
		next := node.Value.(*TreeNode)
		next.Left = nil
		next.Right = nil

		cur.Left = nil
		cur.Right = next

		cur = next
	}
}

/*
将左子树插入到右子树的地方
将原来的右子树接到左子树的最右边节点
考虑新的右子树的根节点，一直重复上边的过程，直到新的右子树为 null

	   1
	  / \
	 2   5
	/ \   \

3   4   6

//将 1 的左子树插入到右子树的地方

	1
	 \
	  2         5
	 / \         \
	3   4         6

//将原来的右子树接到左子树的最右边节点

	   1
	    \
	     2
	    / \
	   3   4
	        \
	         5
	          \
	           6

	//将 2 的左子树插入到右子树的地方
	   1
	    \
	     2
	      \
	       3       4
	                \
	                 5
	                  \
	                   6

	//将原来的右子树接到左子树的最右边节点
	   1
	    \
	     2
	      \
	       3
	        \
	         4
	          \
	           5
	            \
	             6

	 ......

https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/solutions/17274/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--26/
*/
func flatten2(root *TreeNode) {

	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)

		if node.Left != nil {
			pre := node.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			//将右子树挂到 左子树的最右边
			pre.Right = node.Right
			//再将整个左子树挂到根节点的右边
			node.Right = node.Left
			// 不要忘记将node.Left置空
			node.Left = nil
		}

	}
	postOrder(root)
}
