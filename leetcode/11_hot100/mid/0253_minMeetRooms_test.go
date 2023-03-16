package mid

import "sort"

/*
https://leetcode.cn/problems/meeting-rooms-ii/?favorite=2cktkvj

给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，返回 所需会议室的最小数量 。
输入：intervals = [[0,30],[5,10],[15,20]]
输出：2


思路：
从一组想要开会但还没有分配会议室的人员的角度来看这个问题。他们会怎么做？


解题思路
开会也可以理解成坐公交，都是占用某个资源。 就拿题目给的第一组数组来分析。

intervals = [[0,30],[5,10],[15,20]]
第一个人从0上车，从30下车； 第二个人从5上车，10下车。。。

我们的问题转化为最多车上有几个人（也就是最多有多少会议室）。

显然：上车，车上人数+1；下车，车上人数-1 我们把intervals拆解一下

上车：[0, 1], [5, 1], [15, 1]

下车：[10, -1], [20, -1], [30, -1]
然后按照第一个数把上下车排好序

人数 1    2     1     2     1      0
     0----5----10----15----20-----30
变化 +1   +1    -1    +1    -1    -1

作者：muluo
链接：https://leetcode.cn/problems/meeting-rooms-ii/solutions/895579/tu-jie-zhuan-hua-wei-shang-xia-che-wen-t-uy2q/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

这位作者的解法可以理解为，上车就是用会议室，下车就是释放会议室。思路感觉很牛。

*/

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	meetings := make([][]int, 0)
	for _, interval := range intervals {
		meetings = append(meetings, []int{interval[0], 1})  // 上车
		meetings = append(meetings, []int{interval[1], -1}) // 下车
	}
	// 排序
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		} else {
			return meetings[i][0] < meetings[j][0]
		}
	})

	cnt, maxValue := 0, 0
	for _, meeting := range meetings {
		cnt += meeting[1]
		if cnt > maxValue {
			maxValue = cnt
		}
	}
	return maxValue

}
