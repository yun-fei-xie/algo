package _2_linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/description/
给你一个链表的头节点 head 。删除 链表的 中间节点 ，并返回修改后的链表的头节点 head 。
长度为 n 链表的中间节点是从头数起第 ⌊n / 2⌋ 个节点（下标从 0 开始），其中 ⌊x⌋ 表示小于或等于 x 的最大整数。
对于 n = 1、2、3、4 和 5 的情况，中间节点的下标分别是 0、1、1、2 和 2 。

三个指针,pre在中间节点的前面  slow、fast 快慢指针找中间节点
*/
func deleteMiddle(head *ListNode) *ListNode {
	// 只有一个节点或者是空链表的情况下
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = slow.Next
	return head
}

// 这个题如果变一下，1->2->3->4->null 定义2这个位置是中间节点，你怎么求？
// 快慢指针走的过程中，不能让fast走到nil
