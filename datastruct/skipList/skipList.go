package skipList

import (
	"math/rand"
)

/*
leetcode要求对重复数据进行处理。skiplist需要容纳重复数据
*/
const MAXLEVEL = 32
const FACTOR = 0.25

type SkipList struct {
	head  *SkipNode
	level int
	size  int
}

type SkipNode struct {
	key     int
	value   string
	forward []*SkipNode
}

func New() *SkipList {
	return &SkipList{
		head: &SkipNode{
			key:     -1,
			value:   "",
			forward: make([]*SkipNode, MAXLEVEL), //MAXLEVEL个空指针
		},
		level: 0,
	}
}
func (s *SkipList) Size() int {
	return s.size
}
func (s *SkipList) IsEmpty() bool {
	return s.size == 0
}

/*
Get 查找key等于给定key的节点，查不到返回false
*/
func (s *SkipList) Get(key int) (val string, ok bool) {
	current := s.head
	for i := s.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		if current.forward[i] != nil && current.forward[i].key == key {
			return current.forward[i].value, true
		}
	}

	return "", false
}

/*
Remove 删除一个节点，本质也是在查找的过程中记录下前驱节点
*/
func (s *SkipList) Remove(key int) (val string) {
	// 找出节点
	pres := make([]*SkipNode, MAXLEVEL)
	current := s.head
	isExist := false
	for i := s.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		if current.forward[i] != nil && current.forward[i].key == key {
			isExist = true
		}
		pres[i] = current
	}
	if isExist == false { // 节点不存在
		return ""
	}
	node := current.forward[0] // 否则节点就是最底层的current的下一个节点
	// 修改指针，删除节点
	for i := 0; i < s.level && pres[i].forward[i] == node; i++ {

		pres[i].forward[i] = node.forward[i]
		node.forward[i] = nil
	}
	// 更新size
	s.size--

	// 更新level
	for i := s.level; i >= 0; i-- {
		if s.head.forward[i] != nil {
			break
		} else {
			s.level--
		}
	}
	return node.value
}

/*
Put 如果key不存在，则返回(value , true)，如果key已经存在，则返回上一次的v，true
*/
func (s *SkipList) Put(key int, value string) (v string) {
	current := s.head
	pres := make([]*SkipNode, MAXLEVEL)
	for i := 0; i < MAXLEVEL; i++ { // 初始时，前驱都设置为虚拟头结点
		pres[i] = s.head
	}
	for i := s.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		if current.forward[i] != nil && current.forward[i].key == key { // 更新键值对
			lastValue := current.forward[i].value
			current.forward[i].value = value
			return lastValue
		}
		pres[i] = current
	}
	randomLevel := getRandomLevel()
	newNode := &SkipNode{key: key, value: value, forward: make([]*SkipNode, randomLevel)}
	for i := 0; i < randomLevel; i++ {
		newNode.forward[i] = pres[i].forward[i]
		pres[i].forward[i] = newNode
	}
	s.level = max(s.level, randomLevel)
	return value
}

/*
toString 自底向上打印
*/
func (s *SkipList) toString() string {

	// 不太容易将节点对齐，先收集最下面一层到数组，然后用双指针

	return ""
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getRandomLevel() int {
	level := 1
	for rand.Float64() < FACTOR && level < MAXLEVEL {
		level++
	}
	return level
}
