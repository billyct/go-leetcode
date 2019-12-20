package tencent

func LongestPalindrome(s string) string {
	l := len(s)
	// 小于等于 1 个字符的时候，就原样输出
	if l <= 1 {
		return s
	}

	start := 0
	end := 0

	for i := 0; i < l; i++ {
		length1 := expandAroundCenter(&s, i, i)
		length2 := expandAroundCenter(&s, i, i + 1)

		length := length1
		if length1 < length2 {
			length = length2
		}

		if length > end - start + 1 {
			start = i - (length - 1) / 2
			end = i + length / 2
		}
	}

	return s[start: end + 1]
}

// 中心扩散
func expandAroundCenter(s *string, left, right int) int {
	str := *s
	for left >= 0 && right <= len(str) - 1 && str[left] == str[right]{
		left--
		right++
	}

	return right - left - 1
}
