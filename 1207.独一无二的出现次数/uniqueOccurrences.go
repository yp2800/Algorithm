package leetcode

func uniqueOccurrences(arr []int) bool {
	cnts := make(map[int]int)
	for _, v := range arr {
		cnts[v]++
	}

	times := make(map[int]struct{})
	for _, c := range cnts {
		times[c] = struct{}{}
	}
	return len(times) == len(cnts)
}
