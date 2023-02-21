package _2_linkedList

/*
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

一次遍历：快慢指针 slow fast
fast先走n步 然后slow fast 同时走。当fast到达末尾节点时，slow到达需要被删除的节点

因此，fast先走n+1步，这样当fast到达末尾节点时，slow会到达被删除节点的前一个节点 方便删除操作

*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow := dummyHead
	fast := dummyHead

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
