package _2_linkedList

/**
https://leetcode.cn/problems/swap-nodes-in-pairs/
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
输入：head = [1,2,3,4]
输出：[2,1,4,3]

1. 直接交换里面的值
2. 改变指针的方向
3. K个一组

指针交换

```text
A -> B -> C -> D -> null
```
使用递归的方式交换指针的方向。
比如先改变节点C和节点D

1. C指向更改后的子链表
2. D指向C
3. 然后把D返回回去

递归终止条件：尾部只有一个节点或者没有节点的时候，直接将该节点返回（无需处理，解决问题的最小单元）

*/

/*
交换链表中的值
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next
	for first != nil && second != nil {
		temp := first.Val
		first.Val = second.Val
		second.Val = temp

		first = second.Next
		if first == nil {
			break
		}
		second = first.Next
	}
	return head
}

/*
三个指针，p1和p2是需要进行交换的节点，pre负责作为上一段的前驱节点，进行连接。

dummy->1->2->3->4->null
pre    p1 p2

dummy->2->1->3->4->null
pre    p2 p1

dummy->2->1->3->4->null

	pre p1 p2
*/
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	p1 := head
	p2 := p1.Next
	for p1 != nil && p2 != nil {
		// 交换节点
		p1.Next = p2.Next
		p2.Next = p1
		pre.Next = p2
		// 移动节点
		pre = p1
		p1 = p1.Next
		if p1 != nil {
			p2 = p1.Next
		}
	}
	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {

	return swapPairs2Aux(head)

}
func swapPairs2Aux(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(head.Next.Next)
	newHead.Next = head
	return newHead
}
