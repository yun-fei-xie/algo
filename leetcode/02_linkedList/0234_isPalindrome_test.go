package _2_linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode.cn/problems/palindrome-linked-list/description/
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
1->2->2->1->null  中间节点 右边的2
1->2->3->2->1->null  中间节点 中间的3
*/

/*
1->2->2->1->null 能不能通过递归判断
双指针+后序递归
因为后续递归会最先处理最后一个节点，这样可以实现倒着处理链表的节点。也就是从右到左处理节点。
用一个全局变量left,从左到右，在后续递归的时候，插入处理逻辑中，和递归处理的当前节点进行比较。
这种方法其实会把链表各自遍历一遍。
*/
func isPalindrome(head *ListNode) bool {
	left := head

	var dfs func(node *ListNode) bool
	dfs = func(node *ListNode) bool {
		if node == nil {
			return true
		}
		next := dfs(node.Next)
		if next == true && left.Val == node.Val {
			left = left.Next
			return true
		} else {
			return false
		}
	}
	return dfs(head)
}

/*
找到中间节点，然后翻转后半部分
然后用后半部分和前半部分进行对比
并且对比之后恢复链表结构
*/
