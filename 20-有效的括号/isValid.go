package leetcode

func isValid(s string) bool {
	var shed []byte

	reference := make(map[byte]byte)
	reference = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			shed = append(shed, s[i])
			continue
		}
		if len(shed)-1 < 0 {
			return false
		}
		if reference[s[i]] == shed[len(shed)-1] {
			shed = shed[:len(shed)-1]
			continue
		}
		return false
		// fmt.Println(string(value), string(s[0]))
	}
	if len(shed) == 0 {
		return true
	}
	return false
}
