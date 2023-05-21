package main

import (
	"fmt"
	"sort"
)

/*
Description

C 公司对一个即将上线的业务进行了测试，为了更好的评估该业务测试期间的运行状态，C 公司的 Alice 同学打算写一个工具来对该业务的 Access Log 进行处理，她希望能通过这个工具知道任意一个时间段内 (闭区间，精确到秒) 该业务有多少次访问。

为了方便处理，机智的 Alice 同学写了个脚本将 Access Log 简化为了每条日志只有一个时间戳数据，即简化后的 Access Log 是一个 N 行的文本，每行是一个整数时间戳，代表该时刻有一次访问，如下图所示:

1564900123
1564934135
1564934132
1564934666
1564931024
但是 Alice 同学太忙了，所以需要请你帮她实现这个问题。


Input
第 1 行输入 1 个整数
�
T (
1
≤
�
≤
10
1≤T≤10)，表示一共有
�
T 组数据；

对于每一组数据:

第 1 行输入 1 个整数
�
N (
1
≤
�
≤
10000000
1≤N≤10000000)，表示有 Access Log 有
�
N 条数据；
接下来
�
N 行，每行输入一个整数
�
t (
0
≤
�
0≤t，且整数 t 不会超过uint64_t范围)，表示 Access Log 中的一条数据；
接下来 1 行输入 1 个整数 M (
1
≤
�
≤
1000000
1≤M≤1000000)，表示有
�
M 次查询；
接下来
�
M 行，每行两个整数
�
b，
�
e，表示查询
[
�
,
�
]
[b,e] 时间段内的访问次数 (
0
≤
�
≤
�
)
，
并
且
0≤b≤e)，并且b
,
,e$ 不会超过uint64_t范围)。
注意： Access Log 中的时间戳并不能保证是有序的。


Output

*/

func main1() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var N, M int
		fmt.Scan(&N)

		logs := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&logs[j])
		}

		sort.Ints(logs)

		fmt.Scan(&M)
		for j := 0; j < M; j++ {
			var b, e int
			fmt.Scan(&b, &e)

			count := countAccess(logs, b, e)
			fmt.Printf("%d", count)
			if j != M-1 {
				fmt.Println()
			}
		}
	}
}

// [1,2,3,4] [5,6]
func countAccess(logs []int, b, e int) int {
	start := binarySearch(logs, b)
	end := binarySearch(logs, e) - 1

	return end - start
}

func binarySearch(logs []int, t int) int {
	var left, right int
	for left, right = 0, len(logs)-1; left <= right; {
		mid := left + (right-left)/2
		if logs[mid] == t {
			return mid
		} else if logs[mid] > t {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
