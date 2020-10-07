package leetcode

func hammingWeight(num uint32) int {
	power := 31
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		power--
		num = num >> 1
	}
	return count
}

func hammingWeight2(num uint32) int {
	var mask uint32 = 1
	count := 0
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			count++
		}
		mask <<= 1
	}
	return count
}

func hammingWeight3(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num&=(num-1)
	}
	return count
}
