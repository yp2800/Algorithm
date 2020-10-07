package leetcode

func plusOne(digits []int) []int {
	digits[len(digits)-1] += 1
	if digits[len(digits)-1] != 10 {
		return digits
	}
	digits[len(digits)-1] = 0
	carry := 1
	for i := len(digits) - 2; i >= 0; i-- {
		digits[i] += carry
		if digits[i] == 10 {
			digits[i] = 0
			carry = 1
		} else {
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
