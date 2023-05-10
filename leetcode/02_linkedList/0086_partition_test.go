package _2_linkedList

/*
https://leetcode.cn/problems/partition-list/description/

给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。
*/

/*
思路：和快速排序的partition一样 过于复杂 解不出来
思路：维护两个链表 一次遍历head 最后将before的尾巴接在after的头上（跳过虚拟头节点）
*/
func partition(head *ListNode, x int) *ListNode {
	before := &ListNode{Val: -1, Next: nil}
	after := &ListNode{Val: -1, Next: nil}
	tailBefore := before
	tailAfter := after
	cur := head
	for cur != nil {
		if cur.Val < x {
			tailBefore.Next = cur
			cur = cur.Next
			tailBefore = tailBefore.Next
			tailBefore.Next = nil
		} else {
			tailAfter.Next = cur
			cur = cur.Next
			tailAfter = tailAfter.Next
			tailAfter.Next = nil
		}
	}

	tailBefore.Next = after.Next
	return before.Next
}
