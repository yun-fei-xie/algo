# leetcode
## 题目链接

https://leetcode.cn/problems/remove-linked-list-elements/

## 题目描述

给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。


## 解题思路

双指针法：
cur在前，pre紧随其后。
当cur.val == val 时，使用pre将该节点删除。

## 解题代码

```go
type ListNode struct {
Val  int
Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
dummyHead := &ListNode{Next: head}
prev := dummyHead
cur := head

for cur != nil {
if cur.Val == val {
prev.Next = cur.Next
cur = prev.Next
} else {
prev = cur
cur = cur.Next
}
}
return dummyHead.Next

}



```