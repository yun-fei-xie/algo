package bfs_test

import (
	"container/list"
	"fmt"
	"testing"
)

/*
127. 单词接龙
https://leetcode.cn/problems/word-ladder/description/

字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
每一对相邻的单词只差一个字母。
对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
sk == endWord
给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。

方法：构建图，然后使用广度优先遍历
1. 每个不同的单词都看做是图中的一个顶点，如果两个单词只差一个字母，那么这两个顶点之间有搭建了一条边。
2. 然后就是已知起点和终点，求解最短路径的长度。
这个题和433最小基因变化 一模一样

方法：双向广度优先遍历。从起点和终点同时进行遍历。
1. 双向遍历，两个方向各自踏出一步。不断进行。
2.
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var ans int
	var endIndex = -1
	var wordListLen = len(wordList)
	graph := make([][]int, wordListLen)
	for i := 0; i < wordListLen; i++ {
		if wordList[i] == endWord {
			endIndex = i
		}
		for j := i + 1; j < wordListLen; j++ {
			if validChange(wordList[i], wordList[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	if endIndex == -1 {
		return 0
	}

	visited := make([]bool, wordListLen)
	queue := list.New()
	// 将直接与beginWord相连的单词放入队列
	for i := 0; i < wordListLen; i++ {
		if validChange(beginWord, wordList[i]) {
			visited[i] = true
			queue.PushBack(i)
		}
	}

	if queue.Len() == 0 {
		return 0
	}

	// bfs
	step := 1
	for queue.Len() != 0 {
		queueSize := queue.Len()
		for i := 0; i < queueSize; i++ {
			vertex := queue.Remove(queue.Front()).(int)
			if vertex == endIndex {
				// 返回长度
				return step + 1
			}
			for _, adj := range graph[vertex] {
				if !visited[adj] {
					queue.PushBack(adj)
					visited[adj] = true
				}
			}
		}
		step++
	}
	return ans
}

func validChange(gene1 string, gen2 string) bool {
	var cnt int
	for i := 0; i < len(gene1); i++ {
		if gene1[i] != gen2[i] {
			cnt++
		}
	}
	if cnt != 1 {
		return false
	}
	return true
}

/*
双向广度优先遍历
双向广度优先遍历的visited不能使用数组来表示，因为这个时候需要区分节点是之前哪一端访问的。
这里使用hash表。
题目给出了beginWord != endWord 所以，起点和终点不相等
*/
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	var wordListLen = len(wordList)
	var endIndex = -1
	var beginIndex = wordListLen
	graph := make([][]int, wordListLen+1)
	//为beginWord留下一个位置，beginWord编号为0

	for i := 0; i < wordListLen; i++ {
		if wordList[i] == endWord {
			endIndex = i
		}
		//单独处理beginWord和list中的连接关系
		if beginWord != wordList[i] && validChange(beginWord, wordList[i]) {
			graph[beginIndex] = append(graph[beginIndex], i)
			graph[i] = append(graph[i], beginIndex)
		}

		for j := i + 1; j < wordListLen; j++ {
			if validChange(wordList[i], wordList[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	if endIndex == -1 {
		return 0
	}

	visited1 := make(map[int]int) // 单词->访问的距离
	visited2 := make(map[int]int)
	queue1 := list.New()
	queue2 := list.New()
	// 将起点和终点分别放入queue1和queue2
	queue1.PushBack(beginIndex)
	queue2.PushBack(endIndex)
	visited1[beginIndex] = 0
	visited2[endIndex] = 0

	// bfs
	for queue1.Len() != 0 && queue2.Len() != 0 {

		q1Size := queue1.Len()
		q2Size := queue2.Len()
		var step int
		if q1Size < q2Size {
			step = traverse(graph, queue1, visited1, visited2)
		} else {
			step = traverse(graph, queue2, visited2, visited1)
		}
		if step != -1 {
			return step + 1 // 返回单词数目（单词数量=步长+1）
		}

	}
	return 0
}

/*
返回值-1表示还没有到达终点
*/
func traverse(graph [][]int, queue *list.List, visit map[int]int, other map[int]int) int {

	size := queue.Len()
	for i := 0; i < size; i++ {
		vertex := queue.Remove(queue.Front()).(int)

		for _, adj := range graph[vertex] {
			// 如果已经被当前方向访问过
			if _, found := visit[adj]; found {
				continue
			}
			// 如果被对方访问过
			if _, found := other[adj]; found {
				return visit[vertex] + 1 + other[adj]
			}
			queue.PushBack(adj)
			visit[adj] = visit[vertex] + 1
		}
	}
	return -1
}

func TestLadderLength(t *testing.T) {
	fmt.Println(ladderLength2("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength2("a", "c", []string{"a", "b", "c"}))
	fmt.Println(ladderLength("a", "c", []string{"a", "b", "c"}))

}
