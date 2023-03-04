# leetcode

## 题目链接
https://leetcode.cn/problems/swap-nodes-in-pairs/

## 题目描述
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

## 解题思路

指针交换

```text
A -> B -> C -> D -> null
```
使用递归的方式交换指针的方向。
比如先改变节点C和节点D

1. C指向更改后的子链表
2. D指向C
3. 然后把D返回回去

递归终止条件：尾部只有一个节点或者没有节点的时候，直接将该节点返回（无需处理，解决问题的最小单元）

## 解题代码

```go

func swapPairs2(head *ListNode) *ListNode {

	return swapPairs2Aux(head)

}
func swapPairs2Aux(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(head.Next.Next)
	newHead.Next = head
	return newHead
}

```