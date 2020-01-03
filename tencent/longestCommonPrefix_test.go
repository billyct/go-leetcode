package tencent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	as := assert.New(t)

	as.Equal("fl", LongestCommonPrefix([]string{"flower","flow","flight"}))
	as.Equal("", LongestCommonPrefix([]string{"dog","racecar","car"}))
	as.Equal("a", LongestCommonPrefix([]string{"aa", "a"}))
	as.Equal("a", LongestCommonPrefix([]string{"a"}))
}
