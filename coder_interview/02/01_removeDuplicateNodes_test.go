package _2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
*
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

	输入：[1, 2, 3, 3, 2, 1]
	输出：[1, 2, 3]

示例2:

	输入：[1, 1, 1, 1, 2]
	输出：[1, 2]

提示：

链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。
进阶：

如果不得使用临时缓冲区，该怎么解决？
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var record = make(map[int]struct{})
	var prev = head
	var cur = head.Next
	record[prev.Val] = struct{}{}

	// begin check from  second node

	for cur != nil {
		val := cur.Val
		if _, found := record[val]; found {
			prev.Next = cur.Next
			cur = prev.Next
		} else {
			record[cur.Val] = struct{}{}
			cur = cur.Next
			prev = prev.Next
		}
	}

	return head
}

func TestRemoveDuplicateNodes(t *testing.T) {
	data := []int{1, 2, 3, 3, 2, 1}
	listNode, err := makeListNode(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(listNode.toString())
	res := removeDuplicateNodes(listNode)
	fmt.Println(res.toString())

}

func (l *ListNode) toString() string {
	if l == nil {
		fmt.Println("[]")
	}
	sb := strings.Builder{}
	sb.WriteString("[")

	cur := l
	for cur.Next != nil {
		sb.WriteString(strconv.Itoa(cur.Val) + " ,")
		cur = cur.Next
	}
	sb.WriteString(strconv.Itoa(cur.Val) + " ]")
	return sb.String()
}

func makeListNode(arr []int) (*ListNode, error) {
	if len(arr) == 0 {
		return &ListNode{}, errors.New("empty err")
	}
	var head = &ListNode{
		Val:  arr[0],
		Next: nil,
	}
	var tail *ListNode = head

	for i := 1; i < len(arr); i++ {
		tail.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		tail = tail.Next
	}
	return head, nil
}
