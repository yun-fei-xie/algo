package oppo_2023_03_13

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param node1 ListNode类
 * @param node2 ListNode类
 * @return ListNode类
 */
func combineTwoDisorderNodeToOrder(node1 *ListNode, node2 *ListNode) *ListNode {
	// write code here
	arr := make([]int, 0)
	cur1 := node1
	for cur1 != nil {
		arr = append(arr, cur1.Val)
		cur1 = cur1.Next
	}
	cur1 = node2
	for cur1 != nil {
		arr = append(arr, cur1.Val)
		cur1 = cur1.Next
	}

	sort.Ints(arr)

	dummyHead := &ListNode{}
	tail := dummyHead

	for i := 0; i < len(arr); i++ {
		tail.Next = &ListNode{Val: arr[i]}
		tail = tail.Next
	}
	return dummyHead.Next
}
