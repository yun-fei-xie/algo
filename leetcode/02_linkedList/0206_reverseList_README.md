# leetcode

## 题目链接
https://leetcode.cn/problems/reverse-linked-list/
## 题目描述
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

## 解题思路

这个题比较trick的地方在于
第二行代码，两次next的使用，把指针逆转回来
第四行代码，递归链向上抛总是最后那个节点（保证了最末端节点成为新的头）

```go

newHead := reverseListAux(head.Next) 
head.Next.Next = head
head.Next = nil
return newHead // 网上抛的一直是那个头

```



## 解题code


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return reverseListAux(head)
}

func reverseListAux(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}

	newHead := reverseListAux(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead // 网上抛的一直是那个头

```