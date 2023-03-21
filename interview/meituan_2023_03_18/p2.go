package main

import "fmt"

/*

时间限制： 3000MS
内存限制： 589824KB
题目描述：
小美现在有一串彩带，假定每一厘米的彩带上都是一种色彩。

因为任务的需要，小美希望从彩带上截取一段，使得彩带中的颜色数量不超过K种。

显然，这样的截取方法可能非常多。于是小美决定尽量长地截取一段。

你的任务是帮助小美截取尽量长的一段，使得这段彩带上不同的色彩数量不超过K种。



输入描述
第一行两个整数N,K，以空格分开，分别表示彩带有N厘米长，你截取的一段连续的彩带不能超过K种颜色。

接下来一行N个整数，每个整数表示一种色彩，相同的整数表示相同的色彩。

1≤N,K≤5000，彩带上的颜色数字介于[1, 2000]之间。

输出描述
一行，一个整数，表示选取的彩带的最大长度。


8 3
[1,2,3,2,1,4,5,1]

只对了54%


*/

func main2() {
	var n, k int
	fmt.Scanln(&n, &k)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	maxLength := 0
	mp := map[int]int{}
	// 窗口合法性：颜色不超过k种
	for left, right := 0, 0; left <= right && right < len(arr); {
		mp[arr[right]]++ // right

		if len(mp) <= k { // 窗口合法
			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}
			right++

		} else { // 窗口非法-> 让窗口再次合法

			for len(mp) > k {
				mp[arr[left]]--
				if mp[arr[left]] == 0 {
					delete(mp, arr[left])
				}
				left++
			}
		}
	}

	fmt.Println(maxLength)

}
