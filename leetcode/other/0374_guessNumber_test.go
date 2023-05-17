package other

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	return -1
}
func guessNumber(n int) int {

	for left, right := 1, n; left <= right; {
		mid := left + (right-left)/2

		ret := guess(mid)
		if ret == 0 {
			return mid
		} else if ret == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}
