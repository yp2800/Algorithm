package leetcode

func countBits(n int) []int {
	bits := make([]int, n+1)
	f := func(x int) (ones int) {
		for ; x > 0; x &= x - 1 {
			ones++
		}
		return
	}
	for i := range bits {
		bits[i] = f(i)
	}
	return bits
}
