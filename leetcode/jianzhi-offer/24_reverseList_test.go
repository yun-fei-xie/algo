package jianzhi_offer

/*
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL




*/

/*
解法1：双指针迭代
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next // 暂存后继节点
		cur.Next = pre   // 修改指针
		pre = cur        // 移动前驱节点
		cur = next       // 处理下一个节点
	}
	return pre
}

/*
解法2：递归
递归函数的返回值是已经处理完毕的，新的节点的头结点
*/
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var recReverse func(node *ListNode) *ListNode
	recReverse = func(node *ListNode) *ListNode {
		if node.Next == nil {
			return node
		}
		h := recReverse(node.Next)
		node.Next.Next = node
		node.Next = nil
		return h
	}
	return recReverse(head)
}
