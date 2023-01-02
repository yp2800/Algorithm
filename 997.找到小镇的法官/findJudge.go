package leetcode

func findJudge(n int, trust [][]int) int {
	inDegress := make([]int, n+1)
	outDegress := make([]int, n+1)
	for _, t := range trust {
		inDegress[t[1]]++
		outDegress[t[0]]++
	}
	for i := 1; i <= n; i++ {
		if inDegress[i] == n-1 && outDegress[i] == 0 {
			return i
		}
	}
	return -1
}
