package leetcode

func findTheDifference(s string, t string) byte {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; ok {
			m[t[i]]--
		} else {
			return t[i]
		}
	}

	for b, i := range m {
		if i < 0 {
			return b
		}
	}

	return ' '
}

func findTheDifference2(s, t string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i := 0; ; i++ {
		ch := t[i]
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return ch
		}
	}
}

func findTheDifference3(s, t string) byte {
	sum := 0
	for _, ch := range s {
		sum -= int(ch)
	}
	for _, ch := range t {
		sum += int(ch)
	}
	return byte(sum)
}

func findTheDifference4(s, t string) (diff byte) {
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}
