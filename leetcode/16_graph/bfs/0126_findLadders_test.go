package bfs

import (
	"container/list"
	"fmt"
	"testing"
)

/*
126. 单词接龙 II
https://leetcode.cn/problems/word-ladder-ii/description/

按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord -> s1 -> s2 -> ... -> sk 这样的单词序列，并满足：
每对相邻的单词之间仅有单个字母不同。
转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单词。
sk == endWord
给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord 的 最短转换序列 ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。

示例 1：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
解释：存在 2 种最短的转换序列：
"hit" -> "hot" -> "dot" -> "dog" -> "cog"
"hit" -> "hot" -> "lot" -> "log" -> "cog"
示例 2：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：[]
解释：endWord "cog" 不在字典 wordList 中，所以不存在符合要求的转换序列。

在126的基础上，要求给出所有的最短路径
回忆迪杰斯特拉算法，这个问题本质也是一个单源点最短路径问题。
能不能在迪杰斯特拉算法的基础上添加一些信息，回溯求解。

方法：前半部分遍历、遍历的过程需要有一个from结构记住当前单词是从哪些单词过来的。

	后半部分回溯，从from结构中恢复路径。
*/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordDict := make(map[string]struct{})
	for _, word := range wordList {
		wordDict[word] = struct{}{}
	}
	if _, found := wordDict[endWord]; !found {
		return nil
	}
	delete(wordDict, beginWord)

	// 从begin开始搜索

	queue := list.New()
	visited := make(map[string]bool)  // 之前是否已经搜索过
	found := false                    // 是否找到的标志位
	from := make(map[string][]string) // 当前单词可以由哪些单词过来
	visited[beginWord] = true
	queue.PushBack(beginWord)

bfs:
	for queue.Len() != 0 {

		size := queue.Len()
		for i := 0; i < size; i++ {
			if found {
				break bfs
			}
			currentWord := queue.Remove(queue.Front()).(string)
			// 构造单词
			for j := 0; j < len(currentWord); j++ {
				newWord := []byte(currentWord)
				for c := 'a'; c <= 'z'; c++ {
					newWord[j] = byte(c)
					newString := string(newWord)
					// 构造的单词不在wordList中
					if _, found := wordDict[newString]; !found {
						continue
					}

					if currentWord == "lot" {
						if j == 2 && c == 'g' {
							fmt.Println("debug")
						}
					}
					// 构造的单词已经被访问过
					if _, found := visited[newString]; found {
						if newString != endWord {
							continue
						}
					}
					// 单词没有被访问过
					// 但是有一种情况是，终点可以被访问好几次
					from[newString] = append(from[newString], currentWord)
					queue.PushBack(string(newWord))
					visited[string(newWord)] = true
					if string(newWord) == endWord {
						found = true
					}

				}
			}

		}

	}
	var ans = make([][]string, 0)
	// 回溯 构造结果
	var traceBack func(start string, path []string)
	traceBack = func(start string, path []string) {
		if start == beginWord {
			path = append(path, start)
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		path = append(path, start)
		for _, f := range from[start] {
			traceBack(f, path)
		}
	}

	if found {
		path := make([]string, 0)
		traceBack(endWord, path)
	}
	return ans
}

func TestFindLadders(t *testing.T) {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
