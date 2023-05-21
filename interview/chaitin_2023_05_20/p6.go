package main

import "fmt"

func main() {

	var t int
	var k int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&k)
		var ans int
		var num int
		var mod1, mod2 int
		for j := 0; j < k; j++ {
			fmt.Scan(&num)
			// 余数为3
			if num%3 == 0 {
				ans++
			} else if num%3 == 1 {
				// 余数为1
				mod1++
			} else if num%3 == 2 {
				// 余数为2
				mod2++
			}
		}
		fmt.Println(ans + combin(mod1, mod2))
	}
}

func combin(mod1 int, mod2 int) int {
	if mod1 <= mod2 {
		return mod1 + (mod2-mod1)/3
	}
	return (mod2) + (mod1-mod2)/3
}

//func combin(mod1 int, mod2 int) (ans int) {
//	for {
//		if mod2 >= 1 && mod1 >= 1 {
//			ans = ans + 1
//			mod2 = mod2 - 1
//			mod1 = mod1 - 1
//		} else if mod2 == 0 && mod1 >= 3 {
//			mod2 = mod2 + 1
//			mod1 = mod1 - 2
//		} else {
//			break
//		}
//	}
//	return ans
//}
