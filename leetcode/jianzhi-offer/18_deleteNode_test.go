package jianzhi_offer

import (
	"fmt"
	"testing"
)

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	pre := dummyHead
	// 必要条件 pre.Next!=nil  否则pre指向末尾节点会出现pre.next为空，此时执行pre.Next.Val 会引起空指针异常
	for pre != nil {
		if pre.Next != nil && pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return dummyHead.Next
}

func constructListNode(arr []int) *ListNode {
	dummyHead := new(ListNode)
	tail := dummyHead
	for i := 0; i < len(arr); i++ {
		tail.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		tail = tail.Next
	}
	return dummyHead.Next
}
func TestDeleteNode(t *testing.T) {
	head := constructListNode([]int{4, 5, 1, 9})
	l := deleteNode(head, 5)
	fmt.Println(l)
}
