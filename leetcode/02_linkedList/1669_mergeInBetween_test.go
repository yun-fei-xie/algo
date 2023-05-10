package _2_linkedList

/*
https://leetcode.cn/problems/merge-in-between-linked-lists/

给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。

输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
输出：[0,1,2,1000000,1000001,1000002,5]
解释：我们删除 list1 中下标为 3 和 4 的两个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。



*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
方法：模拟
找准4个关键点，然后串起来
1. list1的位置a的前驱节点preA
2. list2的起始节点startList2
3. list2的尾部节点list2Tail
3. list1的位置b的后继节点successorB

把preA和startList2接在一起
把list2Tail和successorB接在一起
*/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	preA := list1
	for i := 0; i < a-1; i++ {
		preA = preA.Next
	}

	successorB := preA
	for j := 0; j < (b - a + 2); j++ {
		successorB = successorB.Next
	}

	preA.Next = list2
	list2Tail := list2
	for list2Tail.Next != nil {
		list2Tail = list2Tail.Next
	}
	list2Tail.Next = successorB

	return list1
}
