package main

import "fmt"

/*
rsa 解密
方法：扩展欧几里得+快速幂
*/
var p = 23333
var q = 10007

func mainP5() {
	//var e, c int = 3, 123612763
	var e, c int
	fmt.Scanf("%d %d", &e, &c)
	decrypt(e, c)
}

func decrypt(e int, c int) {
	phi := (p - 1) * (q - 1)
	d, _ := exgcd(e, phi)
	if d < 0 {
		d = phi + d
	}

	fmt.Println(qpow(uint64(c), uint64(d)))
}

func qpow(a uint64, n uint64) uint64 {
	if n == 0 {
		return 1
	} else if n%2 == 1 {
		return qpow(a, n-1) * a % uint64(p*q)
	} else {
		temp := qpow(a, n/2) % uint64(p*q)
		return temp * temp % uint64(p*q)
	}

}

// 扩展欧几里德算法
func exgcd(a int, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}
	x, y := exgcd(b, a%b)
	x, y = y, x-(a/b)*y
	return x, y
}
