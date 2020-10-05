package tencent

func reverseWords(s string) string {
	// 反转词汇
	rev := func(s []byte, a, b int) {
		for a < b {

			s[a], s[b] = s[b], s[a]

			a++
			b--
		}
	}

	// 最后位置增加空格
	s = s + " "

	ss := []byte(s)
	start := 0

	for i, c := range ss {
		// 当有空格的时候，进行反转词汇
		if c == ' ' {
			rev(ss, start, i - 1)
			start = i + 1
		}
	}

	return string(ss[:len(ss) - 1])
}
