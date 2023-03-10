package _6_stackandqueue

import (
	"container/list"
	"testing"
)

/*
https://leetcode.cn/problems/implement-queue-using-stacks/description/

用栈实现队列，需要思考两个方面，入队、出队。
栈是先入后出、队列是先入先出。
如果使用栈实现先入先出的功能，那么需要使用两个栈。

当需要进行出队操作时，如果第二栈为空，那么将第一个栈中的全部元素逐个pop出来，放入到第二个栈中。
这样经过两次入栈操作，会得到负负得正的效果。
*/

type MyQueue struct {
	pushStack *list.List
	popStack  *list.List
}

/*
构造函数执行后，队列中不应该有元素
*/
func Constructor() MyQueue {
	return MyQueue{
		pushStack: list.New(),
		popStack:  list.New(),
	}
}

func (this *MyQueue) Push(x int) {

	this.pushStack.PushBack(x)

}

func (this *MyQueue) Pop() int {

	if this.popStack.Len() == 0 {
		this.migrateStack()
	}

	element := this.popStack.Back()
	this.popStack.Remove(element)
	return element.Value.(int)

}

func (this *MyQueue) Peek() int {
	if this.popStack.Len() == 0 {
		this.migrateStack()
	}
	element := this.popStack.Back()
	return element.Value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.pushStack.Len() == 0 && this.popStack.Len() == 0
}

func (this *MyQueue) migrateStack() {

	for this.pushStack.Len() != 0 {

		element := this.pushStack.Back()
		this.pushStack.Remove(element)
		this.popStack.PushBack(element.Value) // 这里要把Value取出来 不然会报空指针异常

	}
}

func TestMyQueue(t *testing.T) {

	queue := Constructor()
	queue.Push(1)
	queue.Push(1)
	queue.Peek()
	queue.Pop()
	queue.Empty()
}
