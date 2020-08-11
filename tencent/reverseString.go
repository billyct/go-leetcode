package tencent

func reverseString(s []byte) {
	// 第一种方法
	//left := 0
	//right := len(s) - 1
	//
	//for left < right {
	//	s[left], s[right] = s[right], s[left]
	//	left++
	//	right--
	//}

	// 第二种方法
	//l := len(s)
	//if l <= 1 {
	//	return
	//}
	//for i := 0; i < l/2; i++ {
	//	s[i], s[l-i-1] = s[l-i-1], s[i]
	//}

	// 第三种方法
	reverse(s, 0, len(s) - 1)
}

func reverse(s []byte, a, b int)  {
	if a >= b {
		return
	}

	s[a], s[b] = s[b], s[a]
	reverse(s, a + 1, b - 1)
}
