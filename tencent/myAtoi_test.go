package tencent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	as := assert.New(t)

	as.Equal(42, MyAtoi("42"))
	as.Equal(-42, MyAtoi("   -42"))
	as.Equal(4193, MyAtoi("4193 with words"))
	as.Equal(0, MyAtoi("words and 987"))
	as.Equal(-2147483648, MyAtoi("-91283472332"))
	as.Equal(1, MyAtoi("+1"))
	as.Equal(0, MyAtoi("-+1"))
	as.Equal(2147483647, MyAtoi("2147483648"))
	as.Equal(2147483647, MyAtoi("20000000000000000000"))
	as.Equal(12345678, MyAtoi("  0000000000012345678"))
	as.Equal(0, MyAtoi(""))
	as.Equal(0, MyAtoi("  "))
	as.Equal(0, MyAtoi("+"))
}
