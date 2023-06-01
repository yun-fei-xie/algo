package topSort_test

import (
	"container/list"
	"fmt"
	"testing"
)

func eventualSafeNodes(graph [][]int) []int {
	g := make([][]int, len(graph))
	inDegree := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		g[i] = make([]int, 0)
		// 反向图
		for j := 0; j < len(graph[i]); j++ {
			g[graph[i][j]] = append(g[graph[i][j]], i)
		}
		inDegree[i] = len(graph[i])
	}

	queue := list.New()

	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	for queue.Len() != 0 {
		node := queue.Remove(queue.Front()).(int)

		for _, adj := range g[node] {
			inDegree[adj]--
			if inDegree[adj] == 0 {
				queue.PushBack(adj)
			}
		}
	}

	var ans = make([]int, 0)
	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func TestEventualSafeNodes(t *testing.T) {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
}
