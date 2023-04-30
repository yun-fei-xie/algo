package jianzhi_offer

/*
逆序打印单链表
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var ans = make([]int, 0)
	var dfs func(node *ListNode)
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}

		dfs(node.Next)
		ans = append(ans, node.Val)
	}
	dfs(head)
	return ans
}
