package _2_linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode.cn/problems/remove-nodes-from-linked-list/description/

给你一个链表的头节点 head 。
对于列表中的每个节点 node ，如果其右侧存在一个具有 严格更大 值的节点，则移除 node 。
返回修改后链表的头节点 head 。

输入：head = [5,2,13,3,8]
输出：[13,8]
解释：需要移除的节点是 5 ，2 和 3 。
- 节点 13 在节点 5 右侧。
- 节点 13 在节点 2 右侧。
- 节点 8 在节点 3 右侧。

1 <= Node.val <= 105
*/

/*
解法1：既然要倒着看最大值，那么用递归解决是最合适的，毕竟递归本质就是在倒着遍历链表。
*/
func removeNodes(head *ListNode) *ListNode {
	//记录子链表的头结点和子链表中的最大值
	var dfs func(node *ListNode) (maxVal int, newHead *ListNode)
	dfs = func(node *ListNode) (maxVal int, newHead *ListNode) {
		if node == nil {
			return -1, nil
		}

		maxV, h := dfs(node.Next)
		// 观察发现，子链表的最大值就在子链表的头中，因此不需要用maxVal
		if maxV > node.Val {
			return maxV, h
		} else {
			node.Next = h
			return node.Val, node
		}
	}
	_, newHead := dfs(head)
	return newHead
}

/*
题目保证了节点不为空
*/
func removeNodes2(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	newHead := removeNodes(head.Next)
	if head.Val < newHead.Val {
		return newHead
	} else {
		head.Next = newHead
		return head
	}
}
