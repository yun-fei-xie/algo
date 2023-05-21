package main

import (
	"fmt"
)

/*
Description

公司终于安好饮水机了，Monster 迫不及待要去接水，但是他到那里才发现前面已经有n个同事了。他数了数，饮水机一共有m个接水口。所有的同事严格按照先来后到去接水（m个接水口同时工作，哪个水龙头有空人们就去哪里，如果
�
&lt;
�
n<m，那么就只有n个接水口工作）。每个人都有一个接水的时间，当一个人接完水后，另一个人马上去接，不会浪费时间。Monster 着急要开会，所以他想知道什么时候才能轮到他。

Input
第一行两个整数n和m，表示 Monster 前面有n个人，饮水机有m个接水口。
�
,
�
&lt;
1100
n,m<1100。第二行n个整数，表示每个同学的接水时间。

Output
一行，一个数，表示轮到 monster 接水的时间
*/
func main7() {
	var n, m int // n个同事，m个水龙头
	fmt.Scanln(&n, &m)
	faucet := make([]int, m)
	coWorkers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&coWorkers[i])
	}
	// 不需要等待
	if m > n {
		fmt.Println(0)
		return
	}

	for i := 0; i < m; i++ {
		faucet[i] = coWorkers[i]
	}
	workerPos := m
	var cnt int

	for {
		cnt++

		for i := 0; i < m; i++ {
			faucet[i]--
			if faucet[i] == 0 {
				// 没人了
				if workerPos == n {
					fmt.Println(cnt + 1)
					return
				}
				// 候补上
				faucet[i] = coWorkers[workerPos]
				workerPos++
			}
		}

	}

}
