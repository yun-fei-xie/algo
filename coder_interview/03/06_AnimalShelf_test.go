package _3

import "container/list"

/*
*动物收容所。有家动物收容所只收容狗与猫，且严格遵守“先进先出”的原则。在收养该收容所的动物时，收养人只能收养所有动物中“最老”（由其进入收容所的时间长短而定）的动物，或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。换言之，收养人不能自由挑选想收养的对象。请创建适用于这个系统的数据结构，实现各种操作方法，比如enqueue、dequeueAny、dequeueDog和dequeueCat。允许使用Java内置的LinkedList数据结构。
enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物种类，其中 0 代表猫，1 代表狗。
dequeue*方法返回一个列表[动物编号, 动物种类]，若没有可以收养的动物，则返回[-1,-1]。

思路：两个队列（push进去的信息是动物的编号）当遇到dequeue时，对dog和cat的number进行比较
*/
type AnimalShelf struct {
	cat list.List
	dog list.List
}

func Constructor3() AnimalShelf {
	return AnimalShelf{
		cat: list.List{},
		dog: list.List{},
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	number := animal[0]
	categor := animal[1]
	if categor == 0 {
		this.cat.PushBack(number)
	} else {
		this.dog.PushBack(number)
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	c := this.cat.Front()
	d := this.dog.Front()

	if c == nil && d == nil {
		return []int{-1, -1}
	} else if c == nil && d != nil {
		this.dog.Remove(d)
		return []int{d.Value.(int), 1}
	} else if c != nil && d == nil {
		this.cat.Remove(c)
		return []int{c.Value.(int), 0}
	} else {

		if c.Value.(int) > d.Value.(int) {
			this.dog.Remove(d)
			return []int{d.Value.(int), 1}
		} else {
			this.cat.Remove(c)
			return []int{c.Value.(int), 0}
		}
	}
}
func (this *AnimalShelf) DequeueDog() []int {
	if this.dog.Len() == 0 {
		return []int{-1, -1}
	} else {
		d := this.dog.Front()
		this.dog.Remove(d)
		number := d.Value.(int)
		return []int{number, 1}
	}
}

func (this *AnimalShelf) DequeueCat() []int {
	if this.cat.Len() == 0 {
		return []int{-1, -1}
	} else {
		c := this.cat.Front()
		this.cat.Remove(c)
		number := c.Value.(int)
		return []int{number, 0}
	}

} /**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
