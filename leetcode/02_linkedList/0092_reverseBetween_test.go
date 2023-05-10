package _2_linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

初始时：1-2->3->4->5->null
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	// dummy->1->2->3->4->5->null
	//找到left节点和它的前驱precursor(如果是第一个节点，会没有前驱，使用虚拟头结点)
	p0 := dummyHead
	// 循环left-1次，到达left的前一个节点
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}
	//dummy->1->2->3->4->5->null
	//      p0

	//------
	var pre *ListNode
	var cur = p0.Next
	for i := 0; i < (right-left)+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// 这段代码使用了反转单链表的方法 pre是反转区间新的头，初始为null
	// cur是当前待翻转的节点。区间处理完毕时，cur会走到right节点的下一个位置。
	//  dummy->1 (p0)
	//          \
	//     null<-2<-3<-4    5->null
	//                pre  cur
	//-----

	// ----
	//注意到p0始终指向区间翻转链表的尾部 可以将2指向5 将p0指向4
	p0.Next.Next = cur
	p0.Next = pre
	//----
	return dummyHead.Next
}
