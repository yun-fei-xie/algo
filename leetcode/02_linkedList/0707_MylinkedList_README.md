# leetcode
## 题目链接
https://leetcode.cn/problems/design-linked-list/

## 题目描述

设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

## 解题思路

需要注意的点是，构造函数返回的节点应该作为虚拟头结点
```go
// 这里返回的节点应该作为虚拟头结点
func Constructor() MyLinkedList {
	return MyLinkedList{
		val:  0,
		next: nil,
	}
}

```


## 解题代码


```go
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


```