package tencent

const (
	int32Max = 1<<31 - 1
	int32Min = -1 << 31
)

func MyAtoi(str string) int {
	var i, j int

	n := len(str)
	neg := false

LOOP:
	// 得到第一个数字的节点
	for i = 0; i < n; i++ {
		switch {
		case isNum(str[i]):
			break LOOP
		case str[i] == '-':
			neg = true
			i++
			break LOOP
		case str[i] == '+':
			i++
			break LOOP
		case str[i] != ' ':
			return 0
		}
	}

	ret := 0
	for j = i; j < n; j++ {
		if !isNum(str[j]) {
			break
		}

		num := int(str[j] - '0')
		if neg {
			ret = ret*10 - num
			if ret < int32Min {
				return int32Min
			}

		} else {
			ret = ret*10 + num

			if ret > int32Max {
				return int32Max
			}
		}
	}

	return ret
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}
