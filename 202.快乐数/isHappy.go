package leetcode

// HashSet 检测循环
func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

// 快慢指针法
func isHappy2(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

// 数学法
func isHappy3(n int) bool {
	cycle := map[int]bool{
		4:   true,
		6:   true,
		37:  true,
		58:  true,
		89:  true,
		145: true,
		42:  true,
		20:  true,
	}
	for n != 1 && !cycle[n] {
		n = step(n)
	}
	return n == 1
}
