# leetcode

## 题目链接

https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

## 题目描述

给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

## 解题思路

1. 先求出两个链表的距离之差，然后在作差,用双指针。
2. 有没有一次遍历就解决问题的方法呢？

## 解题代码

move用来移动指针
getLength用于计算链表的长度


```go

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA := getLength(headA)
	lengthB := getLength(headB)

	curA := headA
	curB := headB

	if lengthA > lengthB {
		curA = move(curA, lengthA-lengthB)
	} else {
		curB = move(curB, lengthB-lengthA)
	}

	for curA != nil {

		if curA == curB {
			return curA
		} else {
			curA = curA.Next
			curB = curB.Next
		}

	}
	return nil
}

func move(node *ListNode, step int) *ListNode {

	for i := 0; i < step; i++ {
		node = node.Next
	}
	return node
}

func getLength(l *ListNode) int {
	length := 0
	if l == nil {
		return length
	}
	cur := l
	for cur != nil {
		length++
		cur = cur.Next
	}
	return length
}


```