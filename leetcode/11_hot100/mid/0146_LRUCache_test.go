package mid

/*

https://leetcode.cn/problems/lru-cache/?favorite=2cktkvj

请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；
如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。


思路：使用hash表和双向链表实现


这个代码比较有意思的是在讲一个节点移动到头部的时候，是通过删除+插入实现的，而不是一点一点移动。
这个比较trick

*/
// 双向链表
type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

func initDLinkedNode(key int, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

// LRUcCache
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, found := this.cache[key]; !found {
		return -1
	}
	node, _ := this.cache[key]
	this.moveToHead(node) // 移动到双向链表的头部
	return node.value

}

func (this *LRUCache) Put(key int, value int) {
	if _, found := this.cache[key]; !found {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		// 缓存淘汰 删除末尾的节点
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else { // 将命中的节点放入到双向链表的头部

		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail.pre
	l.removeNode(node)
	return node
}

// 将node移动到双向链表的头部
// 这里的方法是：先将该节点删除，然后再将该节点插入到双向链表的头部
func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (l *LRUCache) addToHead(node *DLinkedNode) {
	node.pre = l.head
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
