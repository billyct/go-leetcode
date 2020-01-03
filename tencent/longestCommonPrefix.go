package tencent

func LongestCommonPrefix(ss []string) string {
	if len(ss) > 0 {
		var ans string

		ans = ss[0]

		for _, s := range ss[1:] {
			i := 0

			for ; i < len(ans) && i < len(s); i++ {
				if ans[i] != s[i] {
					break
				}
			}

			ans = s[:i]

			if ans == "" {
				return ""
			}
		}

		return ans
	}

	return ""
}

