package _2_linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	tail := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			tail = tail.Next
			list1 = list1.Next
			tail.Next = nil
		} else {
			tail.Next = list2
			tail = tail.Next
			list2 = list2.Next
			tail.Next = nil
		}
	}

	if list1 == nil {
		tail.Next = list2
	}

	if list2 == nil {
		tail.Next = list1
	}
	return dummyHead.Next

}
