package tencent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	as := assert.New(t)
	var s string

	s = LongestPalindrome("babad")
	as.Equal("bab", s)

	s = LongestPalindrome("cbbd")
	as.Equal("bb", s)

	s = LongestPalindrome("a")
	as.Equal("a", s)

	s = LongestPalindrome("ac")
	as.Equal("a", s)

	s = LongestPalindrome("bb")
	as.Equal("bb", s)

	s = LongestPalindrome("")
	as.Equal("", s)
}