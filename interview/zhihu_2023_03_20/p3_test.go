package zhihu

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param list1 ListNode类
 * @param list2 ListNode类
 * @return ListNode类
 */

/*
{2,5,1},{1,2,3}

1->5->2
3->2->1

结果 3->7->4
*/
func addTwoNumbers(list1 *ListNode, list2 *ListNode) *ListNode {
	// write code here

	l1, l2 := list1, list2

	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	tail := dummyHead

	nextP := 0
	for l1 != nil && l2 != nil {
		val1 := l1.Val
		val2 := l2.Val

		v := (val1 + val2 + nextP) % 10
		nextP = (val1 + val2 + nextP) / 10

		tail.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		tail = tail.Next

		l1 = l1.Next
		l2 = l2.Next

	}

	if l1 == nil && l2 == nil { // 同时为空
		if nextP != 0 {
			tail.Next = &ListNode{
				Val:  nextP,
				Next: nil,
			}
		}
		return dummyHead.Next
	}
	for l1 != nil {
		v := (l1.Val + nextP) % 10
		nextP = (l1.Val + nextP) / 10
		tail.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		tail = tail.Next
		l1 = l1.Next
	}

	for l2 != nil {
		v := (l2.Val + nextP) % 10
		nextP = (l2.Val + nextP) / 10
		tail.Next = &ListNode{Val: v, Next: nil}
		tail = tail.Next
		l2 = l2.Next
	}

	if nextP != 0 {
		tail.Next = &ListNode{
			Val:  nextP,
			Next: nil,
		}

	}
	return dummyHead.Next
}
