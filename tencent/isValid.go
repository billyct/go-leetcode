package tencent

func IsValid(s string) bool {
	if len(s) & 1 == 1 {
		return false
	}

	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		c := s[i]

		if v, ok := m[c]; ok {
			stack = append(stack, v)
			continue
		}

		if len(stack) > 0 && stack[len(stack) - 1] == c{
			stack = stack[:len(stack) - 1]
			continue
		}

		return false
	}

	return len(stack) == 0
}
