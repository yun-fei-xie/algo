package _3_math

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}

	for n%4 == 0 {
		n = n / 4
	}
	return n == 1
}
