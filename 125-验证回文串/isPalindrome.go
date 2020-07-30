package leetcode

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}
	n := len(sgood)
	sgood = strings.ToLower(sgood)
	for i := 0; i < n/2; i++ {
		if sgood[i] != sgood[n-1-i] {
			return false
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isPalindrome3(s string) bool {
	isValid := func(v rune) bool {
		return unicode.IsDigit(v) || unicode.IsLetter(v)
	}
	left, right := 0, len(s)-1
	for left < right {
		vLeft, vRight := rune(s[left]), rune(s[right])
		if !isValid(vLeft) && !isValid(vRight) {
			left++
			right--
		} else if !isValid(vLeft) {
			left++
		} else if !isValid(vRight) {
			right--
		} else if unicode.ToLower(vLeft) != unicode.ToLower(vRight) {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
