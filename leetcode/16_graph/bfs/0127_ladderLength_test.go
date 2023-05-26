package bfs_test

import "container/list"

/*
127. 单词接龙
https://leetcode.cn/problems/word-ladder/description/

字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
每一对相邻的单词只差一个字母。
对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
sk == endWord
给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0 。

方法：构建图，然后使用广度优先遍历
这个题和433最小基因变化 一模一样
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
