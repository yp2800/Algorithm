package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	return 0
}
func guessNumber(n int) int {
	for i := 0; i < n; i++ {
		if guess(i) == 0 {
			return i
		}
	}
	return n
}

func guessNumber2(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res < 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
