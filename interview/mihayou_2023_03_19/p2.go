package main

import "fmt"

/*

可以进行任意次以下两种操作

每次可以从s删除一个子序列 : mhy
也可以向s中添加一个子序列 : mhy  子序列不需要连续



输入:

3
mhbdy
bd
mhbdy
mhmbhdyy
mhy
abc


输出：
Yes
Yes
No



没什么思路。等等

通过添加和删除操作，最后两个字符串一致
再同一次操作中，应该只会进行添加或者删除中的其中一种。（也有可能既添加 也删除-> 属于调整位置）

知道了。 如果可行，那么两个字符串只差 mhy  的若干倍


发现这类题 不能用模拟的思路去做。
因为它的本质是让你去判断，而不是去得到构造的步骤。

*/

func main() {
	var cnt int
	fmt.Scanln(&cnt)

	var data = make([][2]string, 0)
	for i := 0; i < cnt; i++ {
		var s1 string
		var s2 string
		fmt.Scanln(&s1)
		fmt.Scanln(&s2)
		data = append(data, [2]string{s1, s2})
	}
	var res = make([]bool, 0)
	for i := 0; i < len(data); i++ {
		s1 := data[i][0]
		s2 := data[i][1]

		arr1 := [26]uint8{}
		arr2 := [26]uint8{}

		for j := 0; j < len(s1); j++ {
			arr1[s1[j]-'a']++
		}
		for k := 0; k < len(s2); k++ {
			arr2[s2[k]-'a']++
		}

		// 只能有m h y 这三个位置不同 并且差值必须一致
		var flag = true
		for p := 0; p < len(arr1); p++ {

			if p == ('m'-'a') || p == ('h'-'a') || (p == 'y'-'a') {
				continue
			}
			if arr1[p] != arr2[p] {
				res = append(res, false)
				flag = false
				break
			}
		}
		// 通过了第一关的检验
		if flag {
			if (arr1['m'-'a']-arr2['m'-'a']) == (arr1['h'-'a']-arr2['h'-'a']) && (arr1['h'-'a']-arr2['h'-'a']) == (arr1['y'-'a']-arr2['y'-'a']) {
				res = append(res, true)
			} else {
				res = append(res, false)
			}
		}

	}
	for i := 0; i < len(res); i++ {
		if res[i] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
