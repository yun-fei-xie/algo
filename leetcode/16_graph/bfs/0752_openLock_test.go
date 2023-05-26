package bfs_test

import (
	"container/list"
	"fmt"
	"testing"
)

/*
752. 打开转盘锁
https://leetcode.cn/problems/open-the-lock/description/

你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

方法：广度优先遍历
1. 不需要提前把每一个数字都生成存储好，只需要写一个函数，能将当前状态（数字）生成下一组状态。

方法：A star寻路算法
*/
func openLock(deadends []string, target string) int {
	// 起点就是终点
	if target == "0000" {
		return 0
	}
	// 死亡数字map 方便后续查找
	deadNums := make(map[string]struct{})
	for _, dead := range deadends {
		deadNums[dead] = struct{}{}
	}
	//起点在死亡数字中
	if _, found := deadNums["0000"]; found {
		return -1
	}

	var getNextNums func(num string) []string
	getNextNums = func(num string) []string {
		var ans []string
		bNum := []byte(num)
		for i, b := range bNum {
			bNum[i] = b + 1
			if b+1 > '9' {
				bNum[i] = '0'
			}
			ans = append(ans, string(bNum))
			bNum[i] = b - 1
			if b-1 < '0' {
				bNum[i] = '9'
			}
			ans = append(ans, string(bNum))
			bNum[i] = b
		}
		return ans
	}

	queue := list.New()
	visited := make(map[string]struct{})
	queue.PushBack("0000")
	visited["0000"] = struct{}{}

	step := 0
	for queue.Len() != 0 {
		queueSize := queue.Len()
		for i := 0; i < queueSize; i++ {
			num := queue.Remove(queue.Front()).(string)
			if num == target {
				return step
			}
			nextNums := getNextNums(num)
			for _, next := range nextNums {
				if _, found := deadNums[next]; found {
					continue
				}

				if _, found := visited[next]; found {
					continue
				}
				queue.PushBack(next)
				visited[next] = struct{}{}
			}
		}
		step++
	}
	return -1
}

func TestOpenLock(t *testing.T) {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
