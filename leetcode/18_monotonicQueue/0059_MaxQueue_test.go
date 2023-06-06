package _8_monotonicQueue_test

import "container/list"

/*
剑指 Offer 59 - II. 队列的最大值
https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/description/

请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value 需要返回 -1

方法：单调队列搭配一个普通队列
普通队列用于记录真正的队列中的元素，用于pop出队的时候，判断单调队列的队首元素是否需要出队。

*/

type MaxQueue struct {
	mq *list.List
	q  *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{mq: list.New(), q: list.New()}
}

func (this *MaxQueue) Max_value() int {
	if this.q.Len() == 0 {
		return -1
	}
	return this.mq.Front().Value.(int)

}

func (this *MaxQueue) Push_back(value int) {
	this.q.PushBack(value)

	for this.mq.Len() != 0 && this.mq.Back().Value.(int) < value {
		this.mq.Remove(this.mq.Back())
	}
	this.mq.PushBack(value)

}

func (this *MaxQueue) Pop_front() int {
	if this.q.Len() == 0 {
		return -1
	}
	val := this.q.Front()
	this.q.Remove(this.q.Front())
	if this.mq.Front().Value.(int) == val.Value.(int) {
		this.mq.Remove(this.mq.Front())
	}
	return val.Value.(int)
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
