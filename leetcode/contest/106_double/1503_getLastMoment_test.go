package _06_double

/*
1503. 所有蚂蚁掉下来前的最后一刻
https://leetcode.cn/problems/last-moment-before-all-ants-fall-out-of-a-plank/description/

1.最后一只掉下来的蚂蚁肯定是距离两端最远的哪一只蚂蚁。
2.离两端最远有2层含义：如果是向右，那么计算距离右端的距离；如果是向左，那么计算距离左端的距离
3.具体计算，在向右的数组中找最小值，这个最小值距离右端最远；在向左的数组中找最大值，这个最大值距离左端最远。
*/

func getLastMoment(n int, left []int, right []int) int {
	// 全部都向右走
	if len(left) == 0 {
		return n - min(right...)
	}
	// 全部都向左走
	if len(right) == 0 {
		return max(left...)
	}
	// 一部分向左走 一部分向右走
	leftMax := max(left...)
	rightMin := min(right...)
	return max(leftMax, n-rightMin)
}

func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}
func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}
