package jianzhi_offer

import (
	"fmt"
	"testing"
)

/*
快速幂
*/

func myPow(x float64, n int) float64 {

	var pow func(a float64, b int) float64
	pow = func(a float64, b int) float64 {
		var res float64 = 1.0

		for b > 1 {
			if b%2 == 0 {
				a = a * a
				b = b / 2
			} else {
				res *= a
				a = a * a
				b = b / 2
			}
		}
		return a * res
	}

	if n == 0 {
		return 1
	} else if n > 0 {
		return pow(x, n)
	} else {
		return 1 / pow(x, -n)
	}

}
func TestMyPow(t *testing.T) {
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.000, 10))
	fmt.Println(myPow(2.0, -2))
}
