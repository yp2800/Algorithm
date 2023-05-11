package leetcode

import "fmt"

func countTime(time string) int {
	ans := 0
	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m++ {
			s := fmt.Sprintf("%02d:%02d", h, m)
			ok := 1
			for i := 0; i < 5; i++ {
				if s[i] != time[i] && time[i] != '?' {
					ok = 0
					break
				}
			}
			ans += ok
		}
	}
	return ans
}

func countTime2(time string) int {
	f := func(s string, m int) (cnt int) {
		for i := 0; i < m; i++ {
			a := s[0] == '?' || int(s[0]-'0') == i/10
			b := s[1] == '?' || int(s[1]-'0') == i%10
			if a && b {
				cnt++
			}
		}
		return
	}
	return f(time[:2], 24) * f(time[3:], 60)
}
