package bfs

import (
	"container/list"
	"fmt"
	"testing"
)

/*
433. 最小基因变化
https://leetcode.cn/problems/minimum-genetic-mutation/description/

基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。（变化后的基因必须位于基因库 bank 中）
给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。

方法：广度优先遍历
0.如何抽象出问题？将起始基因看做是起点，终止基因看做是终点，基因库看作是中间节点。那么就是要通过基因库找到一条从起点到终点的最短路径。
1.如何建立图模型？如果两个基因只有一个字母不同，那么这两个基因之间就可以建立一条双向边。给每个基因都编上号。一共有2+len(bank)
2.细节处理？将bank构建成图模型的时候，检查endGene是否在bank中，如果不在直接返回-1。如果在，则记录它在图中的编号。
3.细节处理？start如何与bank之间架设上关系？需要单独处理。将startGene与bank中的基因逐个比对(startGene与bank[i])。如果符合变化一次，则将bank[i]
基因加入队列，并将visited数组标记为true。
*/
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	var endIndex = -1
	graph := make([][]int, len(bank))
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 0)
	}
	// 构建基因库 双向边
	for i := 0; i < len(bank); i++ {
		if bank[i] == endGene {
			endIndex = i
		}
		for j := i + 1; j < len(bank); j++ {
			if validChange(bank[i], bank[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	if endIndex == -1 {
		return -1
	}

	visit := make([]bool, len(bank))
	queue := list.New()
	// 处理startGen
	for i := 0; i < len(bank); i++ {
		if validChange(startGene, bank[i]) {
			visit[i] = true
			queue.PushBack(i)
		}
	}
	if queue.Len() == 0 {
		return -1
	}

	// bfs
	var step = 1
	for queue.Len() != 0 {
		queueSize := queue.Len()
		for i := 0; i < queueSize; i++ {

			vertex := queue.Remove(queue.Front()).(int)
			if vertex == endIndex {
				return step
			}
			for _, adj := range graph[vertex] {
				if !visit[adj] {
					queue.PushBack(adj)
					visit[adj] = true
				}
			}
		}
		step++
	}

	return -1
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

func TestMinMutation(t *testing.T) {
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}
