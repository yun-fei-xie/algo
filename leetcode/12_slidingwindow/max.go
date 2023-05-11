package _2_slidingwindow

func max(num ...int) int {
	m := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > m {
			m = num[i]
		}
	}
	return m
}

func min(num ...int) int {
	m := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] < m {
			m = num[i]
		}
	}
	return m
}
