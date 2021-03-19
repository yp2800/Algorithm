package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
