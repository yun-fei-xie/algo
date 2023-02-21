package _2_linkedList

/**
https://leetcode.cn/problems/reverse-linked-list/
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
用一个数组这种思路就不写了
递归翻转：把指针的方向调整过来
注意的是，需要记住最终头节点
https://leetcode.cn/problems/reverse-linked-list/solutions/551596/fan-zhuan-lian-biao-by-leetcode-solution-d1k2/
官方题解讲解的清楚
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return reverseListAux(head)
}

func reverseListAux(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}

	newHead := reverseListAux(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead // 网上抛的一直是那个头
}
