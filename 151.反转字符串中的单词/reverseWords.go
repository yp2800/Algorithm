package leetcode

import (
	"sort"
	"strings"
)

func reverseWords(s string) string {
	ans := ""
	split := strings.Split(s, " ")
	for _, s2 := range split {
		if s2 == "" {
			continue
		}
		if ans != "" {
			s2 = s2 + " "
		}
		ans = s2 + ans
	}
	return ans
}

func reverseWords2(s string) string {
	split := strings.Split(s, " ")
	sort.Slice(split, func(i, j int) bool {
		return true
	})

	ans := make([]string, 0)
	for i := range split {
		if split[i] == "" {
			continue
		}
		ans = append(ans, split[i])
	}
	return strings.Join(ans, " ")
}

func reverseWords3(s string) string {
	// 移除空白字符
	ss := []byte(trim_spaces(s))
	//翻转字符串
	reverse(ss, 0, len(ss)-1)
	//翻转每一个单词
	reverse_each_word(ss)

	return string(ss)
}

func reverse_each_word(ss []byte) {
	n := len(ss)
	start, end := 0, 0
	for start < n {
		for end < n && ss[end] != ' ' {
			end++
		}
		reverse(ss, start, end-1)
		start = end + 1
		end++
	}
}

func reverse(ss []byte, left, right int) {
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left, right = left+1, right-1
	}
}
func reverse2(ss []byte) {
	var a []byte
	for _, s := range ss {
		defer func(v byte) {
			a = append(a, v)
		}(s)
	}
}

func trim_spaces(s string) string {
	//移除左侧空白字符
	left, right := 0, len(s)-1
	for left <= right && s[left] == ' ' {
		left++
	}
	//移除右侧空白字符
	for left <= right && s[right] == ' ' {
		right--
	}
	//	移除中间空白字符
	output := make([]byte, 0)
	for left <= right {
		if s[left] != ' ' {
			output = append(output, s[left])
		} else if output[len(output)-1] != ' ' {
			output = append(output, s[left])
		}
		left++
	}
	return string(output)
}
