package _2_linkedList

/*
*
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }

https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

先求出两个链表的距离之差
然后在作差 用双指针
*/
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

// 代码尽量这样写 划分成不同的函数
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
