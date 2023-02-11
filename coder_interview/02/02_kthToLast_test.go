package _2

import (
	"fmt"
	"testing"
)

/**
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。

思路：如果ListNode长度为l ， 那么倒数第k个节点的index为 l-k  index范围[0 , l-1]

需要走index步
*/

func kthToLast(head *ListNode, k int) int {
	length := getLength(head)
	index := length - k
	cur := head

	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func TestKthToLast(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 7, 8}
	listNode, _ := makeListNode(arr)

	res := kthToLast(listNode, 3)
	fmt.Println(res)

}

func getLength(l *ListNode) int {
	length := 0
	if l == nil {
		return length
	}
	cur := l
	for cur != nil {
		length++
		cur = cur.Next
	}
	return length
}
