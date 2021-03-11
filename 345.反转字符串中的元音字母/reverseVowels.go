package leetcode

import (
	"log"
	"reflect"
	"unsafe"
)

func reverseVowels(s string) string {
	// 类型转换，会发生内在拷备
	b := []byte(s)
	m := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	for i, j := 0, len(s)-1; i < j; {
		for _, ok := m[s[i]]; !ok && i < j; _, ok = m[s[i]] {
			i++
		}
		for _, ok := m[s[j]]; !ok && i < j; _, ok = m[s[j]] {
			j--
			log.Println("J", j)
		}
		if i < j {
			b[i], b[j] = s[j], s[i]
			j--
			i++
		}
	}
	log.Println(s)
	return string(b)
}

func reverseVowels2(s string) string {

	// 类型转换，会发生内在拷备
	//b:=[]byte(s)
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: ssh.Data,
		Len:  ssh.Len,
		Cap:  ssh.Len,
	}
	b := *(*[]byte)(unsafe.Pointer(&bh))
	m := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	for i, j := 0, len(b)-1; i < j; {
		for _, ok := m[b[i]]; !ok && i < j; _, ok = m[b[i]] {
			i++
		}
		for _, ok := m[b[j]]; !ok && i < j; _, ok = m[b[j]] {
			j--
			log.Println("J", j)
		}
		if i < j {
			b[i], b[j] = b[j], b[i]
			j--
			i++
		}
	}
	return string(b)
}

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
