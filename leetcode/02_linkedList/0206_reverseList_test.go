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
方法1：用一个数组这种思路就不写了
方法2：递归翻转：把指针的方向调整过来。需要注意的是，需要记住最终头节点。参考官方题解：https://leetcode.cn/problems/reverse-linked-list/solutions/551596/fan-zhuan-lian-biao-by-leetcode-solution-d1k2/
方法3：迭代翻转
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

/*
头插法翻转
1->2->3->4-null
-----
dummyHead->1->null
dummyHead->2->1->null
dummyHead->3->2->1->null
return dummyHead.next
*/

func reverseList2(head *ListNode) *ListNode {
	dummyHead := new(ListNode)
	cur := head
	for cur != nil {
		// 记住原始链表的下一个节点
		next := cur.Next
		//将待插入节点断开
		cur.Next = nil
		// 在dummyHead尾部，其他节点的头部 插入节点
		cur.Next = dummyHead.Next
		dummyHead.Next = cur
		//cur回到原始链表
		cur = next
	}

	return dummyHead.Next
}

/*
在原始链表上迭代翻转
使用两个指针
newHead代表新的链表的头
cur代表需要翻转的节点
初始状态

		null     1->2->3->4->null

		 ^       ^

	  newHead   cur
*/
func reverseList3(head *ListNode) *ListNode {
	var newHead *ListNode //
	var cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = newHead // 重新建立关系
		newHead = cur
		cur = next
	}
	return newHead
}
