package _2_linkedList

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/merge-k-sorted-lists/

给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[

	1->4->5,
	1->3->4,
	2->6

]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]
*/

/*
思路一：将链表分组 两两合并 调用21题的mergeTwoList
思路二：使用优先队列
*/
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	// 递归二分->类似归并排序
	var conquer func(left int, right int) *ListNode
	conquer = func(left int, right int) *ListNode {
		if left > right {
			return nil
		}
		// 一条链表
		if left == right {
			return lists[left]
		}
		// 有两条链表
		if left+1 == right {
			return mergeTwoLists(lists[left], lists[right])
		}

		mid := left + (right-left)/2
		leftMerge := conquer(left, mid)
		rightMerge := conquer(mid+1, right)
		return mergeTwoList(leftMerge, rightMerge)
	}
	return conquer(0, length-1)
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {

	dummyHead := &ListNode{
		Val:  -1,
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
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return dummyHead.Next
}

func TestMergeKList(t *testing.T) {
	//[[],[1]]
	lists := make([]*ListNode, 0, 2)
	lists = append(lists, nil)

	list2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	lists = append(lists, list2)

	ans := mergeKLists(lists)
	fmt.Println(ans)

}
