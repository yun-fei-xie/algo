package _3

import "math"

/**
请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(1)。
*/

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  []int{math.MaxInt},
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	m := this.GetMin()
	this.min = append(this.min, min(x, m))
}

// 用length指着 （怎么释放空间） slice 应该不会浪费太多性能
func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
