package _45

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/find-the-losers-of-the-circular-game/
6430 找出转圈优秀的输家

n 个朋友在玩游戏。这些朋友坐成一个圈，按 顺时针方向 从 1 到 n 编号。从第 i 个朋友的位置开始顺时针移动 1 步会到达第 (i + 1) 个朋友的位置（1 <= i < n），而从第 n 个朋友的位置开始顺时针移动 1 步会回到第 1 个朋友的位置。

游戏规则如下：

第 1 个朋友接球。

接着，第 1 个朋友将球传给距离他顺时针方向 k 步的朋友。
然后，接球的朋友应该把球传给距离他顺时针方向 2 * k 步的朋友。
接着，接球的朋友应该把球传给距离他顺时针方向 3 * k 步的朋友，以此类推。
换句话说，在第 i 轮中持有球的那位朋友需要将球传递给距离他顺时针方向 i * k 步的朋友。

当某个朋友第 2 次接到球时，游戏结束。

在整场游戏中没有接到过球的朋友是 输家 。

给你参与游戏的朋友数量 n 和一个整数 k ，请按升序排列返回包含所有输家编号的数组 answer 作为答案。

0k、1k、2k
1-n

编号平移到[0...n-1]方便取模操作
*/
func circularGameLosers(n int, k int) []int {
	visit := make(map[int]bool)

	i := 0
	step := i * k
	pos := 0
	for {
		if visit[pos] == true {
			break
		}
		// 标记
		visit[pos] = true

		// 算出下一个位置
		i++
		step = i * k
		pos = (pos + step) % n
	}

	var ans = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if visit[i] == false {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func TestCircularGameLosers(t *testing.T) {
	fmt.Println(circularGameLosers(5, 2))
	fmt.Println(circularGameLosers(4, 4))
}
