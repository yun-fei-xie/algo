package _2_linkedList

/**
https://leetcode.cn/problems/swap-nodes-in-pairs/

1. 直接交换里面的值
2. 改变指针的方向

*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next
	for first != nil && second != nil {
		temp := first.Val
		first.Val = second.Val
		second.Val = temp

		first = second.Next
		if first == nil {
			break
		}
		second = first.Next
	}
	return head
}

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
