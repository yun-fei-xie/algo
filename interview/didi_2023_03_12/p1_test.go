package didi_2023_03_12

import "fmt"

/*
时间限制： 3000MS
内存限制： 589824KB
题目描述：

	小明正在进行积木的分销。他一共有N个积木，他要将它们分装到M个小包装内，每个小包装内至少有一个。如果一个小包装内含x个积木，那么这个小包装将会被定价为X2。小明想要控制一下价格，不希望价格太贵或者太便宜。他想要知道是否存在一种分装方案，使得分装后的M个小包装定价之和恰好为P。

	如果有多种方案，输出字典序最小的那一个。对于两种不同方案{a1,a2,...,aM}与{b1,b2,...,bM}，若对于1≤i≤t的i均有ai=bi，且at+1＜bt+1，那么认为方案a的字典序小于方案b。

注意：当t=0时，没有合法的i存在，1≤i≤t只是限制i的范围。

例如，对于M=3,N=4的情况下，{1,1,2}的字典序小于{2,1,1} (对应t=0的情况) 、{1,2,1} (对应t=1的情况)。

输入描述
第一行三个正整数N,M,P，含义如题面。

对于所有数据，1≤M≤N≤12；0≤P≤109

输出描述
若不存在任何方案，输出-1，否则输出M个数表示每个小包装内应分的的积木数量
*/
func main1() {
	var n, m, p int // n个积木、m个包、价格总和p  每个包至少有一个 每个包的价格为 x^2
	fmt.Scan(&n, &m, &p)
	// 从小数量开始尝试
	var dfs func(cnt int, bagIndex int, totalPrice int)
	var path = make([]int, 0) // [0...m-1]
	var res = make([][]int, 0)
	dfs = func(cnt int, bagIndex int, totalPrice int) {
		if bagIndex >= m {
			if cnt == 0 && totalPrice == p {
				temp := make([]int, len(path))
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}

		// 当前这个包最少放1个， (当前这个包后面还有 m-bagIndex-1 个包)
		// 最多放 cnt - (m-bagIndex-1)

		for i := 1; i <= cnt-(m-bagIndex-1); i++ {
			path = append(path, i)
			dfs(cnt-i, bagIndex+1, totalPrice+i*i)
			path = path[:len(path)-1]
		}

	}
	dfs(n, 0, 0)
	if len(res) == 0 {
		fmt.Println("-1")
	} else {
		r := res[0]
		for i := 0; i < len(r); i++ {
			fmt.Printf("%d", r[i])
			if i != len(r)-1 {
				fmt.Printf(" ")
			}
		}
	}
}
