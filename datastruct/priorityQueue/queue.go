package priorityQueue

/*
队列的接口
*/
type Queue interface {
	GetSize() int
	IsEmpty() bool
	EnQueue(item any)
	DeQueue() any
	GetFront() any
}
