package _105_double_test

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

/*
6464.最大公约数遍历
https://leetcode.cn/problems/greatest-common-divisor-traversal/

方法：暴力求解 时间空间都爆炸 极端情况下会nums中的所有数字都一样，肯定每个数字都能和其他数字联通。

	这样就是一副完全图。使用邻接表存储，会占用 n^2的空间。如果数组的长度是1万，那么最后的存储差不多就1亿条边。

方法：建图技巧-> 使用公约数作为中转站。

	题目给出了数字的最大值是10^5 也就是10000。可以用两个数字的最大公约数作为中转站。
	这样可以降低边的数量。
	在方法1的邻接表存储过程中，graph[i] = {1,2,3,5} graph[i][j]表示nums[i]和nums[j]之间可达

方法：并查集。如果nums[i]可以被分解为一系列质因数相乘，nums[j]也可以被分解为一系列质因数相乘，如果两者有相同的

	公共质因数，那么将公共质因数相乘就是两者的最大公约数。如果nums[i]和nums[j]有公共的质因数，那意味着两者的最大
	公约数肯定不为1。也就意味着两者可达。使用并查集，将nums[i]和它们的质因子进行合并。
	最后验证，所有的nums[i]和nums[i+1]是否在同一个集合中。(i>=0 && i<len(nums))

1.如何对一个数字分解质因数？
*/
func canTraverseAllPairs(nums []int) bool {
	graph := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if gcd(nums[i], nums[j]) > 1 {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	// 是不是可以再遍历的时候再生成边

	// 任意两个顶点i,j(i>j)是否可达

	var bfs func(start int) bool
	bfs = func(start int) bool {
		queue := list.New()
		visited := make([]bool, len(nums))
		queue.PushBack(0)
		visited[0] = true

		for queue.Len() != 0 {
			queueSize := queue.Len()
			for i := 0; i < queueSize; i++ {
				vertex := queue.Remove(queue.Front()).(int)
				for _, adj := range graph[vertex] {
					if !visited[adj] {
						queue.PushBack(adj)
						visited[adj] = true
					}
				}
			}
		}

		for i := 0; i < len(visited); i++ {
			if !visited[i] {
				return false
			}
		}
		return true
	}
	return bfs(0)
}

func gcd(i, j int) int {
	if j == 0 {
		return i
	} else {
		return gcd(j, i%j)
	}
}

// 使用埃式筛求出10000以内的素数
// 合数标记为true(因为bool数组默认值是false,不想初始化)
func eulerPrimeFilter(n int) []bool {
	table := make([]bool, n+1)
	for i := 0; i < len(table); i++ {
		table[i] = true
	}
	table[0], table[1] = false, false
	table[2] = true
	for i := 2; i < len(table); i++ {
		if !table[i] {
			for j := i << 1; j < len(table); j = j + i {
				table[j] = false
			}
		}
	}
	return table
}

// 朴素方法分解质因数
func breakPrimeFactor(n int) (ans []int) {
	upper := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < upper; i++ {
		for n%i == 0 {
			n = n / i
			ans = append(ans, i)
		}
	}
	if n != 1 {
		ans = append(ans, n)
	}
	return ans
}

func canTraverseAllPairs2(nums []int) bool {
	uf := initUnionFind(100000)
	for i := 0; i < len(nums); i++ {
		pf := breakPrimeFactor(nums[i])
		for j := 0; j < len(pf); j++ {
			uf.unionElements(nums[i], pf[j])
		}
	}

	// 查找
	for i := 0; i < len(nums)-1; i++ {
		if !uf.isConnected(nums[i], nums[i+1]) || (nums[i] == 1 && nums[i+1] == 1) {
			return false
		}
	}
	return true
}

type unionFind struct {
	parent []int
	count  int
}

func initUnionFind(n int) *unionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 初始时，每一个节点的父节点都是本身
	}
	return &unionFind{
		parent: parent,
		count:  n,
	}
}

func (un *unionFind) unionElements(p, q int) {
	pRoot := un.find(p)
	qRoot := un.find(q)
	if pRoot == qRoot {
		return
	} else {
		un.parent[pRoot] = qRoot
	}
}

func (un *unionFind) isConnected(p, q int) bool {
	pRoot := un.find(p)
	qRoot := un.find(q)
	return pRoot == qRoot
}

/*
在查找的过程中实现路径压缩
*/
func (un *unionFind) find(p int) int {
	parentId := un.parent[p]
	if parentId == p {
		return parentId
	}
	un.parent[p] = un.find(un.parent[p])
	return un.parent[p]
}

func TestCanTraverseAllPairs(t *testing.T) {
	fmt.Println(canTraverseAllPairs([]int{4, 3, 12, 8}))
	fmt.Println(canTraverseAllPairs2([]int{4, 3, 12, 8}))
	fmt.Println(canTraverseAllPairs([]int{3, 9, 5}))
	fmt.Println(canTraverseAllPairs2([]int{3, 9, 5}))
	fmt.Println(canTraverseAllPairs([]int{2, 3, 6}))
	fmt.Println(canTraverseAllPairs2([]int{2, 3, 6}))
	//fmt.Println(gcd(4, 3))
	//fmt.Println(gcd(3, 12))
	//fmt.Println(breakPrimeFactor(9))
	//fmt.Println(breakPrimeFactor(1))
	//fmt.Println(canTraverseAllPairs2([]int{1, 1}))
}
