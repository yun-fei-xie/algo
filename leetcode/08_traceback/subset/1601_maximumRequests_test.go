package subset

import (
	"fmt"
	"testing"
)

/*
1601. 最多可达成的换楼请求数目
https://leetcode.cn/problems/maximum-number-of-achievable-transfer-requests/

我们有 n 栋楼，编号从 0 到 n - 1 。每栋楼有若干员工。由于现在是换楼的季节，部分员工想要换一栋楼居住。
给你一个数组 requests ，其中 requests[i] = [fromi, toi] ，表示一个员工请求从编号为 fromi 的楼搬到编号为 toi 的楼。
一开始 所有楼都是满的，所以从请求列表中选出的若干个请求是可行的需要满足 每栋楼员工净变化为 0 。意思是每栋楼 离开 的员工数目 等于 该楼 搬入 的员工数数目。比方说 n = 3 且两个员工要离开楼 0 ，一个员工要离开楼 1 ，一个员工要离开楼 2 ，如果该请求列表可行，应该要有两个员工搬入楼 0 ，一个员工搬入楼 1 ，一个员工搬入楼 2 。
请你从原请求列表中选出若干个请求，使得它们是一个可行的请求列表，并返回所有可行列表中最大请求数目。

输入：n = 5, requests = [[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]]
输出：5
解释：请求列表如下：
从楼 0 离开的员工为 x 和 y ，且他们都想要搬到楼 1 。
从楼 1 离开的员工为 a 和 b ，且他们分别想要搬到楼 2 和 0 。
从楼 2 离开的员工为 z ，且他想要搬到楼 0 。
从楼 3 离开的员工为 c ，且他想要搬到楼 4 。
没有员工从楼 4 离开。
我们可以让 x 和 b 交换他们的楼，以满足他们的请求。
我们可以让 y，a 和 z 三人在三栋楼间交换位置，满足他们的要求。
所以最多可以满足 5 个请求。

方法：枚举。
每一个请求都可以被加入处理列表。进行综合考虑。
当得到一个待考虑列表后，如何判断是否可行？
考虑[3,4] 3号楼少一个人，4号楼多一个人。
由于开始，楼是满的，所有，交换后，所有楼的净流入人口为0。
可以用一个数组进行验证。[3,4] arr[3]-- , arr[4]++
最后，arr[i]必须都等于0
*/

func maximumRequests(n int, requests [][]int) int {
	// 存放requests的下标
	var ans = 0
	var length = len(requests)
	var candidate = make([]int, 0)
	var traceback func(i int)
	traceback = func(i int) {
		if i == length {
			if isValidate(candidate, requests) {
				if len(candidate) > ans {
					ans = len(candidate)
				}
			}
			return
		}
		// 加入
		candidate = append(candidate, i)
		traceback(i + 1)
		// 不加入
		candidate = candidate[0 : len(candidate)-1]
		traceback(i + 1)

	}
	traceback(0)
	return ans
}

func isValidate(candidate []int, requests [][]int) bool {
	var arr = [20]int{}
	for i := 0; i < len(candidate); i++ {
		out := requests[candidate[i]][0]
		in := requests[candidate[i]][1]
		arr[in]++
		arr[out]--
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

func TestMaximumRequests(t *testing.T) {
	//fmt.Println(maximumRequests(4, [][]int{{0, 3}, {3, 1}, {1, 2}, {2, 0}}))
	fmt.Println(maximumRequests(5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}))
}
