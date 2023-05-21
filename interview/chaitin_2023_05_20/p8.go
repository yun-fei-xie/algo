package main

import (
	"fmt"
	"math/big"
)

/*
ac

*/

func main8() {

	var n int
	for {
		cnt, _ := fmt.Scanln(&n)
		if cnt != 0 {
			fmt.Println(fib(n))
		} else {
			break
		}
	}

}

func fib(n int) *big.Int {

	if n == 1 || n == 2 {
		return big.NewInt(1)
	}

	var f1, f2 = big.NewInt(1), big.NewInt(1)
	var f *big.Int
	for i := 3; i <= n; i++ {
		f = f1.Add(f1, f2)
		f1 = f2
		f2 = f
	}
	return f
}
