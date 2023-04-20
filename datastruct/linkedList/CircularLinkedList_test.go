package linkedList

import (
	"fmt"
	"testing"
)

/*
单项循环链表
*/

var head *CircularLinkedNode // 用一个head变量保存循环单项链表的头节点

type CircularLinkedNode struct {
	val  int
	next *CircularLinkedNode
}

/*
初始化返回一个空
*/
func initCircularLinkedList() *CircularLinkedNode {
	return nil
}

/*
将节点插入循环单向链表的尾节点
*/
func (l *CircularLinkedNode) insertNode(val int) {
	node := &CircularLinkedNode{
		val:  val,
		next: nil,
	}

	if head == nil {
		node.next = node
		head = node //初始化head节点
	} else { // 通过遍历找出尾节点，尾节点的next是head
		cur := head
		for cur.next != head {
			cur = cur.next
		}
		cur.next = node
		node.next = head
	}
}

func (l *CircularLinkedNode) print() {
	if head == nil {
		fmt.Println("null")
	} else {
		cur := head
		for cur.next != head {
			fmt.Printf("%d ->", cur.val)
			cur = cur.next
		}
		fmt.Printf("%d -> end\n", cur.val)
	}

}

func TestCircularLinkedList(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	circularLinkedList := initCircularLinkedList()
	for _, num := range arr {
		circularLinkedList.insertNode(num)
	}
	circularLinkedList.print()
}
