package leetcode

func countPrimes(n int) int {
	count := 0
	sings := make(map[int]bool, n)
	for i := 2; i < n; i++ {
		if sings[i] {
			continue
		}
		count++
		for j := 2 * i; j < n; j += i {
			sings[j] = true
		}
	}
	return count
}
