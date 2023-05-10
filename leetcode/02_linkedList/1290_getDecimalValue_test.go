package _2_linkedList

/*
https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/description/
给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。
请你返回该链表所表示数字的 十进制值 。
输入：head = [1,0,1]
输出：5
解释：二进制数 (101) 转化为十进制数 (5)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var ans int
	for ; head != nil; head = head.Next {
		ans = ans*2 + head.Val
	}
	return ans
}
