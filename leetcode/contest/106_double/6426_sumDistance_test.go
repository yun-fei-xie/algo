package _06_double

import (
	"fmt"
	"sort"
	"testing"
)

/*
6426.移动的机器人
https://leetcode.cn/problems/movement-of-robots/description/

对于一个机器人，考虑每一步移动是否会和周围机器人
通过观察可以发现，所有的机器人之间的相对位置是不会改变的。
nums = [-2,0,2], s = "RLL", d = 3

比赛时候的想法，使用模拟，考虑相撞之后的运动情况。十分复杂，因为需要考虑相撞后两个机器人会对周围的机器人造成影响。
（或者真的可以模拟一下）

这个题的思路是，碰撞之后可以认为两个机器人交换了身份。继续向着原来的轨迹继续前进。
这样来整个问题就简单多了。直接将原来的位置按照运动的方向走上d步。

最后便是求每一个机器人距离其他机器人的相对距离。
两两距离可以规约为从左到右每个机器人与它左边的所有机器人的距离之和。
这一步如果暴力计算会超时，通过列式子可以发现使用前缀和可以优化计算。

这个问题的思路和周赛196蚂蚁问题比较相似
https://leetcode.cn/problems/last-moment-before-all-ants-fall-out-of-a-plank/
*/
func sumDistance(nums []int, s string, d int) int {
	mod := 1000000000 + 7
	pos := make([]int, len(nums))
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			pos[i] = nums[i] - d
		} else {
			pos[i] = nums[i] + d
		}
	}
	sort.Ints(pos)
	// 统计和 从第二个机器人开始统计
	var ans int
	for i := 1; i < len(pos); i++ {
		for j := 0; j < i; j++ {
			ans += pos[i] - pos[j]
			ans = ans % mod
		}
	}
	// 对于pos[i]到[0...i-1]中的各个点的距离，(pos[i]-pos[0]) + (pos[i]-pos[1]) + ...+ (pos[i]-pos[i-1])
	// 把每一项的pos[i]提起来得到 （i)*pos[i] - (pos[0]+...+pos[i-1]) 后面就是前缀和。这里有i项。
	// 把这个优化写到第二版代码中
	return ans
}

func sumDistance2(nums []int, s string, d int) int {
	mod := 1000000000 + 7
	pos := make([]int, len(nums))
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			pos[i] = nums[i] - d
		} else {
			pos[i] = nums[i] + d
		}
	}
	sort.Ints(pos)
	// 统计和 从第二个机器人开始统计
	var ans int
	prefix := make([]int, len(nums))
	prefix[0] = pos[0]
	preSum := pos[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = preSum + pos[i]
		preSum = prefix[i]
		ans += i*pos[i] - prefix[i-1]
		ans = ans % mod
	}
	return ans
}

func TestSumDistance(t *testing.T) {
	fmt.Println(sumDistance2([]int{-2, 0, 2}, "RLL", 3))
	fmt.Println(sumDistance2([]int{1, 0}, "RL", 2))
}
