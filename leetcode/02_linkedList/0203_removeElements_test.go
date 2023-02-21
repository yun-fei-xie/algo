package _2_linkedList

/**

 */

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
