package leetcode

// 动态规划
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	//初始化结婚（最小的回文就是单个字符）
	maxLen := 1
	begin := 0

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		//dp[i][j] 表示 s[i..j]是否是回文串
		dp[i][i] = true
	}

	//递推开始
	//先枚举子串长度
	for j := 1; j < n; j++ {
		//枚举左边界，左边界的上限设置可以宽松一些
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

// 中心扩展算法
func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		if right1-left1 > end-start {
			start, end = left1, right1
		}

		left2, right2 := expandAroundCenter(s, i, i+1)
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
