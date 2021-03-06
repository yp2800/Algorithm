package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	ch2Word := map[byte]string{}
	word2Ch := map[string]byte{}
	split := strings.Split(s, " ")

	if len(split) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		word := split[i]
		if word2Ch[word] > 0 && word2Ch[word] != ch || ch2Word[ch] != "" && ch2Word[ch] != word {
			return false
		}
		ch2Word[ch] = word
		word2Ch[word] = ch

	}
	return true
}
