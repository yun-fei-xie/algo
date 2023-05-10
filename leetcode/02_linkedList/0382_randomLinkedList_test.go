package _2_linkedList

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode.cn/problems/linked-list-random-node/
给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。
实现 Solution 类：

Solution(ListNode head) 使用整数数组初始化对象。
int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。

*/

/*
解法1：用一个数组记录链表元素，每次调用就随机返回一个元素。
随机函数直接使用库函数。 rand.Intn()

解法2：水塘抽样,从链表头开始，遍历整个链表，对遍历到的第 iii 个节点，随机选择区间 [0,i)内的一个整数，如果其等于0，则将答案置为该节点值，否则答案不变。

*/

//type Solution struct {
//	arr []int
//}
//
//func Constructor382(head *ListNode) Solution {
//	arr := make([]int, 0)
//	for head != nil {
//		arr = append(arr, head.Val)
//		head = head.Next
//	}
//	return Solution{arr: arr}
//
//}
//
//func (this *Solution) GetRandom() int {
//	length := len(this.arr)
//	randIndex := rand.Intn(length)
//	return this.arr[randIndex]
//}

type Solution struct {
	head *ListNode
}

func Constructor382(head *ListNode) Solution {
	return Solution{head: head}
}

func (this *Solution) GetRandom() int {
	var ans = this.head.Val
	for cur, i := this.head, 1; cur != nil; cur = cur.Next {
		if rand.Intn(i) == 0 {
			ans = cur.Val
		}
		i++
	}
	return ans
}
