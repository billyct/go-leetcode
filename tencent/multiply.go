package tencent

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	l1 := len(num1)
	l2 := len(num2)

	length := l1 + l2

	ans := make([]byte, length)

	var n1, tmp byte

	// 模拟乘法计算
	for i := l1 - 1; i >= 0; i-- {
		// 如果是 0 则继续
		n1 = num1[i] - '0'
		if n1 == 0 {
			continue
		}

		for j := l2 - 1; j >= 0; j-- {
			tmp = (num2[j]-'0')*n1 + ans[i+j+1]
			// 更新当前位数字
			ans[i+j+1] = tmp % 10
			// 更新进位数字
			ans[i+j] += tmp / 10
		}
	}

	s := ""

	for k, v := range ans {
		if k == 0 && v == 0 {
			continue
		}

		s += string(v + '0')
	}

	return s
}
