package unionFind_test

import (
	"container/list"
	"math"
)

/*
952. 按公因数计算最大组件大小
https://leetcode.cn/problems/largest-component-size-by-common-factor/

给定一个由不同正整数的组成的非空数组 nums ，考虑下面的图：
有 nums.length 个节点，按从 nums[0] 到 nums[nums.length - 1] 标记；
只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j]之间才有一条边。
返回 图中最大连通组件的大小 。

方法：求解连通分量中，顶点个数最多的那个那个连通分量的顶点个数.
直接对nums[i]和nums[j]两两求gcd然后建图连边，最后bfs统计，最后会超时。

方法：并查集+质因数分解
将每一个nums[i]进行质因数分解，然后用并查集将nums[i]和nums[i]的质因数合并。
最后遍历nums数组，统计每一组的最大值。
*/
func largestComponentSize(nums []int) int {
	length := len(nums)
	graph := make([][]int, length)
	for i := 0; i < length; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if gcd(nums[i], nums[j]) > 1 {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	// bfs遍历
	visited := make([]bool, length)

	var bfs func(i int) int
	bfs = func(i int) int {
		queue := list.New()
		queue.PushBack(i)
		visited[i] = true
		var count int
		for queue.Len() != 0 {
			size := queue.Len()
			count += size
			for i := 0; i < size; i++ {
				vertex := queue.Remove(queue.Front()).(int)
				for _, adj := range graph[vertex] {
					if !visited[adj] {
						visited[adj] = true
						queue.PushBack(adj)
					}
				}
			}
		}
		return count
	}

	var ans int
	for i := 0; i < length; i++ {
		if !visited[i] {
			count := bfs(i)
			if count > ans {
				ans = count
			}
		}
	}
	return ans
}

func gcd(i, j int) int {
	if j == 0 {
		return i
	}
	return gcd(j, i%j)
}

/*
使用并查集求解
*/
func largestComponentSize2(nums []int) int {
	uf := InitUnionFind2(100000 + 1)
	length := len(nums)
	for i := 0; i < length; i++ {
		factors := breakPrimeFactor(nums[i])
		for _, factor := range factors {
			uf.unionElements(nums[i], factor)
		}
	}

	// 查找每一个root下面有多少孩子
	// 用一个map记录
	counts := make(map[int]int)
	for i := 0; i < length; i++ {
		root := uf.find(nums[i])
		counts[root]++
	}
	// 遍历一次map,拿到value中的最大值
	var ans int
	for _, value := range counts {
		if value > ans {
			ans = value
		}
	}
	return ans
}

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

type unionFind struct {
	parent []int
	count  int
}

func InitUnionFind2(n int) *unionFind {
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

func (un *unionFind) find(p int) int {
	parentId := un.parent[p]
	if parentId == p {
		return parentId
	}
	un.parent[p] = un.find(un.parent[p])
	return un.parent[p]
}

/* 路径压缩
5->4->3->2->1->0
如果查找5的父亲节点，一次find之后，所有的节点都会指向节点0。
*/
