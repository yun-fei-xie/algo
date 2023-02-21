package _2_linkedList

/**
https://leetcode.cn/problems/linked-list-cycle-ii/

解法：只需要检测当前节点的 next有没有被访问过
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	record := make(map[*ListNode]struct{})
	cur := head

	record[cur] = struct{}{}

	for cur.Next != nil {
		if _, found := record[cur.Next]; found {
			return cur.Next
		} else {
			cur = cur.Next
			record[cur] = struct{}{}
		}
	}

	return nil
}
