# leetcode

## 题目链接

https://leetcode.cn/problems/linked-list-cycle-ii/

## 题目描述
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

## 解题思路

1. 用一个hash表记录访问过的节点，指针一直向前移动。 当指针再次访问到hash表中的节点时，则代表链表有环。
2. 如何使用O(1)空间，不额外使用空间？

## 解题代码

```go

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	record := make(map[*ListNode]struct{})
	cur := head

	record[cur] = struct{}{}

	for cur.Next != nil {
		if _, found := record[cur.Next]; found {
			return cur.Next
		} else {
			cur = cur.Next
			record[cur] = struct{}{}
		}
	}

	return nil

```