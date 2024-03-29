package main

import "fmt"

/*
时间限制： 3000MS
内存限制： 589824KB
题目描述：
现在小美获得了一个字符串。小美想要使得这个字符串是回文串。

小美找到了你。你可以将字符串中至多两个位置改为任意小写英文字符’a’-‘z’。

你的任务是帮助小美在当前制约下，获得字典序最小的回文字符串。

数据保证能在题目限制下形成回文字符串。

注：回文字符串：即一个字符串从前向后和从后向前是完全一致的字符串。

例如字符串abcba, aaaa, acca都是回文字符串。字符串abcd, acea都不是回文字符串


和下一个排列有点像。。

肯定能回文！这个很重要


*/

func main3() {

	var s string
	fmt.Scanln(&s)

	left, right := 0, len(s)-1
	arr := []byte(s)

	var maxOp = 2

	for left < right {
		if arr[left] == arr[right] {
			left++
			right--
		} else {
			if arr[left] < arr[right] {
				arr[right] = arr[left]
			} else {
				arr[left] = arr[right]
			}
			maxOp--
		}
	}

	// 这里肯定已经回文
	if maxOp != 0 {
		// 还有调换的机会
		for left, right := 0, len(arr)-1; left < right; {
			if arr[left] != 'a' {
				arr[left] = 'a'
				arr[right] = 'a'
				maxOp -= 2
			}
			left++
			right--
		}

	}

	fmt.Println(string(arr))

}
