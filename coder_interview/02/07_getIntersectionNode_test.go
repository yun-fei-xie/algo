package _2

/*
*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
图示两个链表在节点 c1 开始相交：

题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。
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

func move(node *ListNode, step int) *ListNode {

	for i := 0; i < step; i++ {
		node = node.Next
	}
	return node
}
