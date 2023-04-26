package skipList

import (
	"fmt"
	"testing"
)

func TestSkipList2(t *testing.T) {
	sl := New()
	v0, ok := sl.Get(1)
	if ok {
		fmt.Println(v0)
	}

	sl.Put(1, "hello")
	v1, ok := sl.Get(1)
	if ok {
		fmt.Println(v1)
	}
	sl.Put(2, "world")
	v2, ok := sl.Get(2)
	if ok {
		fmt.Println(v2)
	}
	sl.Put(3, "hello world")
	v3, ok := sl.Get(3)
	if ok {
		fmt.Println(v3)
	}

	sl.Remove(3)
	v4, ok := sl.Get(3)
	if !ok {
		fmt.Println("找不到")
	} else {
		fmt.Println(v4)
	}

}
func TestSkipList3(t *testing.T) {
	// ["Skiplist","add","add","add","add","search","erase","search","search","search"]
	// [[],[0],[5],[2],[1],[0],[5],[2],[3],[2]]
	s := New()
	s.Put(0, "0")
	s.Put(5, "5")
	s.Put(2, "2")
	s.Put(1, "1")

	v0, ok := s.Get(0)
	if ok {
		fmt.Println(v0)
	} else {
		fmt.Println("not found")
	}

	s.Remove(5)
	s.Get(2)
	s.Get(3)
	s.Get(2)

}
