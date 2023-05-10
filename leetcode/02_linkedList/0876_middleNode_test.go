package _2_linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
快慢指针
这次写居然忘了细节
*/
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	//for slow.Next != nil {
	//
	//	slow = slow.Next
	//	if fast.Next != nil && fast.Next.Next != nil {
	//		fast = fast.Next.Next
	//	} else {
	//		return slow
	//	}
	//
	//}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow

}
