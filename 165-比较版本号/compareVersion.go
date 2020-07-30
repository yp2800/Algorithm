package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	version1_array := strings.Split(version1, ".")
	version2_array := strings.Split(version2, ".")

	length := len(version1_array)
	if length < len(version2_array) {
		length = len(version2_array)
		tmp := make([]string, length-len(version1_array))
		for i := range tmp {
			tmp[i] = "0"
		}
		version1_array = append(version1_array, tmp...)
	}
	if length > len(version2_array) {
		tmp := make([]string, length-len(version2_array))
		for i := range tmp {
			tmp[i] = "0"
		}
		version2_array = append(version2_array, tmp...)
	}

	for i := 0; i < length; i++ {
		v1, _ := strconv.Atoi(version1_array[i])
		v2, _ := strconv.Atoi(version2_array[i])
		if v1 < v2 {
			return -1
		}
		if v1 > v2 {
			return 1
		}
	}
	return 0
}
