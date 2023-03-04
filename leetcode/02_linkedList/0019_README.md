# leetcode

## 题目链接
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

## 题目描述
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

## 解题思路

一次遍历：快慢指针 slow fast
fast先走n步 然后slow fast 同时走。当fast到达末尾节点时，slow到达需要被删除的节点

fast先走n+1步，这样当fast到达末尾节点时，slow会到达被删除节点的前一个节点，这样可以方便删除操作。

## 解题代码


```go
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
```