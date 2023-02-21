package _2_linkedList

/**
https://leetcode.cn/problems/design-linked-list/
*/

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

// 这里返回的节点应该作为虚拟头结点
func Constructor() MyLinkedList {
	return MyLinkedList{
		val:  0,
		next: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.next

	for i := 0; i < index; i++ { // i表示从当前位置移动的步数
		if cur != nil {
			cur = cur.next
		} else {
			return -1
		}
	}

	if cur == nil {
		return -1
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyLinkedList{
		val:  val,
		next: nil,
	}
	node.next = this.next
	this.next = node
}

func (this *MyLinkedList) AddAtTail(val int) {

	pre := this
	for pre.next != nil {
		pre = pre.next
	}
	pre.next = &MyLinkedList{
		val:  val,
		next: nil,
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		node := &MyLinkedList{
			val:  val,
			next: nil,
		}
		node.next = this.next
		this.next = node
	}

	prev := this
	var i = 0

	for ; i < index; i++ {
		if prev.next != nil {
			prev = prev.next
		} else {
			break
		}
	}
	// prev.next == nil
	if prev.next == nil && i == index {
		prev.next = &MyLinkedList{val: val}
	} else if prev.next != nil {
		node := &MyLinkedList{val: val, next: prev.next}
		prev.next = node
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	prev := this
	var i int = 0
	for ; i < index; i++ {
		if prev.next != nil {
			prev = prev.next
		} else {
			break
		}
	}

	if prev.next != nil {
		prev.next = prev.next.next
	}
}
