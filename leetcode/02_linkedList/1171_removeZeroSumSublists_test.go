package _2_linkedList

/*
https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/description/

给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。

删除完毕后，请你返回最终结果链表的头节点。



你可以返回任何满足题目要求的答案。

（注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）

示例 1：

输入：head = [1,2,-3,3,1]
输出：[3,1]
提示：答案 [1,2,1] 也是正确的。
示例 2：

输入：head = [1,2,3,-3,4]
输出：[1,2,4]



head = [1,2,-3,3,1]
prefix=[1,3, 0,3,4]

*/

/*
*
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }

// 如果是求总和是任意值的区间呢？

/*
hash

head = [1,2,-3,3,1]
prefix=[1,3, 0,3,4]

利用前缀和，如果sum[0...i]的前缀和与sum[0...j]相同，则sum[i+1...j] = 0
*/
func removeZeroSumSublists2(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	prefixMap := make(map[int]*ListNode)
	prefixMap[0] = dummy
	// 相同前缀会被覆盖
	for cur, sum := head, 0; cur != nil; cur = cur.Next {
		sum += cur.Val
		prefixMap[sum] = cur
	}

	for cur, sum := head, 0; cur != nil; cur = cur.Next {
		sum += cur.Val
		// 这个地方相当巧妙
		if node, found := prefixMap[sum]; found {
			cur.Next = node.Next
		}
	}
	return dummy.Next
}
