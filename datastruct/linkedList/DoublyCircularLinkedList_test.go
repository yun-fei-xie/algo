package linkedList

import (
	"fmt"
	"testing"
)

/*
双向循环链表
*/
var doublyHead *doublyCircularNode

type doublyCircularNode struct {
	val  int
	prev *doublyCircularNode
	next *doublyCircularNode
}

func initDoublyCircularLinkedList() *doublyCircularNode {
	return nil
}

/*
将节点插入到双向循环链表的尾部
node的下一个节点是head
*/
func (l *doublyCircularNode) insertNode(val int) {
	node := &doublyCircularNode{
		val:  val,
		prev: nil,
		next: nil,
	}

	if doublyHead == nil { // 孤家寡人
		doublyHead = node
		doublyHead.prev = doublyHead
		doublyHead.next = doublyHead
	} else {
		cur := doublyHead
		for cur.next != doublyHead {
			cur = cur.next
		}
		cur.next = node
		node.prev = cur
		node.next = doublyHead
		doublyHead.prev = node
	}
}

func (l *doublyCircularNode) print() {
	if doublyHead == nil {
		fmt.Printf("%s\n", "null")
	} else {
		cur := doublyHead
		for cur.next != doublyHead {
			fmt.Printf("%d->", cur.val)
			cur = cur.next
		}
		fmt.Printf("%d->end\n", cur.val)
	}
}

func TestDoubleCircularLinkedList(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	doublyCircularLinkedList := initDoublyCircularLinkedList()
	for _, num := range arr {
		doublyCircularLinkedList.insertNode(num)
	}

	doublyCircularLinkedList.print()

}
