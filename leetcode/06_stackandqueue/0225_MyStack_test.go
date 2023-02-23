package _6_stackandqueue

import (
	"container/list"
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/implement-stack-using-queues/
两个队列实现一个栈

出栈操作如何实现（这个和用栈实现队列还是有区别的）在用栈实现队列的时候，两个栈是输入和输出的关系
但是这里两个队列则不是，其中一个队列用于给另外一个队列进行备份。即出栈的时候，

其实可以让两个队列轮换身份。
*/

type MyStack struct {
	enQueue *list.List
	deQueue *list.List
}

func ConstructorMyStack() MyStack {
	return MyStack{
		enQueue: list.New(),
		deQueue: list.New(),
	}
}

func (this *MyStack) Push(x int) {

	this.enQueue.PushBack(x)

}

func (this *MyStack) Pop() int {

	for this.enQueue.Len() > 1 {
		element := this.enQueue.Front()
		this.enQueue.Remove(element)
		this.deQueue.PushBack(element.Value)
	}
	res := this.enQueue.Front()
	this.enQueue.Remove(res)

	this.enQueue, this.deQueue = this.deQueue, this.enQueue

	return res.Value.(int)

}

func (this *MyStack) Top() int {
	for this.enQueue.Len() > 1 {
		element := this.enQueue.Front()
		this.enQueue.Remove(element)
		this.deQueue.PushBack(element.Value)
	}
	res := this.enQueue.Front()
	this.enQueue.Remove(res)
	this.deQueue.PushBack(res.Value)

	this.enQueue, this.deQueue = this.deQueue, this.enQueue
	return res.Value.(int)
}

func (this *MyStack) Empty() bool {

	return this.enQueue.Len() == 0 && this.deQueue.Len() == 0
}

func TestMyStack(t *testing.T) {
	s := ConstructorMyStack()
	s.Push(1)
	s.Push(2)
	val1 := s.Top()
	val2 := s.Pop()
	isEmpty := s.Empty()
	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println(isEmpty)
}
