package leetcode

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	var r, s string
	var count int
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	s = countAndSay(n - 1)
	//初始化字符计数器
	count = 1
	//1. 对S从左到右遍历
	for i := 1; i <= len(s)-1; i++ {
		//1.1 如果s[i] != s[i-1] 则把前面的字符打印
		if s[i] != s[i-1] {
			r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i-1]))
			count = 1
		} else {
			//1.2 如果相等则计数器+1
			count++
		}
		//1.2 如果到末尾了，则打印当前字符
		if i+1 == len(s) {
			r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i]))
			s = r
			return s
		}
	}
	return s
}

func countAndSay2(n int) string {
	str := "1"
	for i := 2; i <= n; i++ {
		var tmp strings.Builder
		preByte := str[0]
		count := 1
		for j := 1; j < len(str); j++ {
			if str[j] == preByte {
				count++
			} else {
				tmp.WriteString(strconv.Itoa(count))
				tmp.WriteByte(preByte)
				preByte = str[j]
				count = 1
			}
		}
		tmp.WriteString(strconv.Itoa(count))
		tmp.WriteByte(preByte)
		str = tmp.String()
	}
	return str
}

func countAndSay3(n int) string {
	if n == 1 {
		return "1"
	} else {
		last := countAndSay3(n-1)
		preByte := last[0]
		count := 1
		var buffer bytes.Buffer
		for i:= 1;i<len(last);i++{
			if preByte == last[i] {
				count++
			} else {
				buffer.WriteString(strconv.Itoa(count))
				buffer.WriteByte(preByte)
				preByte = last[i]
				count = 1
			}
		}
		buffer.WriteString(strconv.Itoa(count))
		buffer.WriteByte(preByte)
		return buffer.String()
	}
}