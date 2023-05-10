package _2_linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/description/
剑指 Offer 22. 链表中倒数第k个节点

给定一个链表: 1->2->3->4->5, 和 k = 2.
返回链表 4->5.

方法1：统计链表长度length,倒数第k个节点就是整数第（length-k）+1 个节点。设置一个指针指向头结点，然后走（length-k）步。（因为初始指向head就消耗掉一步）
方法2：双指针pre,after初始都指向head,后指针先走k步，然后前后指针一起走(先移动pre,再移动after顺序不能变)。当后指针走向空的时候，前指针指向倒数第k个节点。
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {

	cur := head
	var length int
	for cur != nil {
		length++
		cur = cur.Next
	}

	index := length - k //从0开始计数
	ans := head
	for i := 0; i < index; i++ {
		ans = ans.Next
	}
	return ans
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	pre, after := head, head
	for i := 0; i < k; i++ {
		after = after.Next
	}

	for after != nil {
		pre = pre.Next
		after = after.Next
	}
	return pre
}
