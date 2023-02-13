package _3

import (
	"fmt"
	"testing"
)

/**
实现一个MyQueue类，该类用两个栈来实现一个队列。

思路：一个输入stack , 一个输出stack
当入队时，直接append到输入stack中,
当出队或者peek时，将输入stack中的数据反序列push到输出stack中，然后最末尾的数据返回回去。
所以比较关键的一步就是倒腾数据 in2out
*/

type MyQueue struct {
	inStack, outStack []int
}

/** Initialize your data structure here. */
func Constructor2() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.in2out()
	}
	x := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.in2out()
	}
	return this.outStack[len(this.outStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

func (this *MyQueue) in2out() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func TestMyQueue(t *testing.T) {

	obj := Constructor2()
	obj.Push(2)
	obj.Push(3)
	obj.Push(5)

	p := obj.Peek()
	fmt.Println("peek value ", p)
	r := obj.Pop()
	v := obj.Peek()
	fmt.Printf("pop value %d ,now peek value %d", r, v)

}
