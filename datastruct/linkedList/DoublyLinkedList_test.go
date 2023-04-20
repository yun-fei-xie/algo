package linkedList

import (
	"fmt"
	"testing"
)

/*
双向链表
需要注意的是，双向链表不是循环链表。
它的最后一个元素的next不是链表的第一个元素，而是null。
*/

type DoublyLinkedNode struct {
	value int
	prev  *DoublyLinkedNode
	next  *DoublyLinkedNode
}

/*
init函数返回一个虚拟头结点
*/
func InitDoublyLinkedList() *DoublyLinkedNode {
	return &DoublyLinkedNode{
		value: -1,
		prev:  nil,
		next:  nil,
	}
}

/*
向链表的尾部插入一个元素
*/
func (l *DoublyLinkedNode) InsertTail(v int) {
	cur := l
	for cur.next != nil {
		cur = cur.next
	}
	node := &DoublyLinkedNode{
		value: v,
		prev:  cur,
		next:  nil,
	}
	cur.next = node
}

func (l *DoublyLinkedNode) print() {
	cur := l.next
	for cur != nil {
		fmt.Printf("%d -> ", cur.value)
		cur = cur.next
	}
	fmt.Printf("%s\n", "null")
}

func TestDoublyLinkedList(t *testing.T) {
	doubleList := InitDoublyLinkedList()

	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		doubleList.InsertTail(arr[i])
	}
	doubleList.print()
}
