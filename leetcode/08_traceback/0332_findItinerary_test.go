package _8_traceback

import (
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode.cn/problems/reconstruct-itinerary/

输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
输出：["JFK","MUC","LHR","SFO","SJC"]


输入：tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。

1. 确定起点
2. 把所有元素都
3. 相当于用一根线把连通图串起来
*/

func findItinerary(tickets [][]string) []string {
	res := make([]string, 0)
	used := make([]int, len(tickets))

	var dfs func(tickets [][]string, startIndex int, depth int)
	dfs = func(tickets [][]string, startIndex int, depth int) {
		if depth == len(tickets) {
			res = append(res, tickets[startIndex][1])
			return
		}
		newStart := tickets[startIndex][1] // 新的开头  需要搜索index(可能出现多个)
		newIndex := -1
		first := true
		// 寻找新的起点
		for i := 0; i < len(tickets); i++ {
			if used[i] == 0 && tickets[i][0] == newStart {
				if first {
					newIndex = i
					first = false
				} else {
					last := tickets[newIndex][1]
					cur := tickets[i][1]
					c := strings.Compare(last, cur)
					if c > 0 { //发现字母序更小的
						newIndex = i
					}
				}
			}
		}

		res = append(res, tickets[newIndex][0])
		used[newIndex] = 1
		dfs(tickets, newIndex, depth+1)
	}

	first := true
	startIndex := -1
	for i := 0; i < len(tickets); i++ {
		if tickets[i][0] == "JFK" {
			if first {
				startIndex = i
				first = false
			} else {
				last := tickets[startIndex][1]
				cur := tickets[i][i]
				if c := strings.Compare(last, cur); c > 0 {
					startIndex = i
				}
			}
		}
	}
	used[startIndex] = 1
	res = append(res, "JFK")
	dfs(tickets, startIndex, 1)
	return res
}

func TestFindItinerary(t *testing.T) {
	//tickets := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	//res := findItinerary(tickets)
	//fmt.Println(res)
	//
	//tickets2 := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	//res2 := findItinerary(tickets2)
	//fmt.Println(res2)

	tickets3 := [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}
	res3 := findItinerary(tickets3)
	fmt.Println(res3)
}
