package leetcode

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for num != 1 {
		if num%2 == 0 {
			num = num >> 1
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}
	return true
}
