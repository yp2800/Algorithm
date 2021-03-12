package leetcode

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	left, right := 2, num/2
	for left <= right {
		x := left + (right-left)/2
		guess_squared := x * x
		if guess_squared == num {
			return true
		}
		if guess_squared > num {
			right = x - 1
		} else {
			left = x + 1
		}
	}
	return false
}

func isPerfectSquare2(num int) bool {
	r := num
	for r*r > num {
		r = (r + num/r) / 2
	}
	if r * r == num {
		return true
	}

	return false
}
