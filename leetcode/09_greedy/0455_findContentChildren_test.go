package _9_greedy

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/assign-cookies/description/
每次用当前最大的饼干 喂最大胃口的孩子，如果当前的饼干满足不了当前的孩子，则选择下一个胃口较小的孩子
*/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	i := len(g) - 1 //content
	j := len(s) - 1 // cake

	for i >= 0 && j >= 0 {
		if s[j] >= g[i] { //当前饼干j可以满足i的胃口
			res++
			j--
			i--
		} else { // 当前饼干j无法满足i的胃口
			i--
		}
	}

	return res
}

func TestFindContentChildren(t *testing.T) {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	res := findContentChildren(g, s)
	fmt.Println(res)

}
