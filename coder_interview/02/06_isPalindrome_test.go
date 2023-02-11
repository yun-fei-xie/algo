package _2

/**
编写一个函数，检查输入的链表是否是回文的。
示例 1：

输入： 1->2
输出： false
示例 2：

输入： 1->2->2->1
输出： true
*/
/*
由于无法从后向前遍历-> 可以先把值copy出来判断
如果是奇数个节点 最终tail和head会相遇 相遇时满足回文条件
如果是偶数个节点 最后head.next = tail
*/
func isPalindrome(head *ListNode) bool {
	listNodeValues := getListValues(head)
	start := 0
	end := len(listNodeValues) - 1

	for start <= end {

		if listNodeValues[start] != listNodeValues[end] {
			return false
		}
		start++
		end--

	}
	return true
}

func getListValues(head *ListNode) []int {

	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func getTailNode(node *ListNode) *ListNode {

	if node.Next == nil {
		return node
	}
	for node.Next != nil {
		node = node.Next
	}
	return node

}
