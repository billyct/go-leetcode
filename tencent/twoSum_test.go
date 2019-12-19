package tencent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	as := assert.New(t)
	var result []int

	result = TwoSum([]int{2, 7, 11, 15}, 9)
	as.Equal([]int{0, 1}, result)

	result = TwoSum([]int{3, 2, 4}, 6)
	as.Equal([]int{1, 2}, result)

	result = TwoSum([]int{3, 3}, 6)
	as.Equal([]int{0, 1}, result)

	result = TwoSum([]int{-3,4,3,90}, 0)
	as.Equal([]int{0, 2}, result)
}
