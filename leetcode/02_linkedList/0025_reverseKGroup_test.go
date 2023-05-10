package _2_linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
*/
func reverseKGroup(head *ListNode, k int) *ListNode {

	var length int
	h := head
	for h != nil {
		length++
		h = h.Next
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	p0 := dummy
	var pre *ListNode
	var cur = p0.Next
	//k个一组翻转不足一组不翻转
	for length/k != 0 {
		length = length - k

		// 翻转一组
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}

		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		pre = nil
		p0 = nxt

	}

	return dummy.Next
}
