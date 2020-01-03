package tencent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	as := assert.New(t)

	as.Equal([][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}, ThreeSum([]int{-1, 0, 1, 2, -1, -4}))

	as.Equal([][]int{{0,0,0}}, ThreeSum([]int{0,0,0}))
	as.Equal([][]int{{-1,0,1}}, ThreeSum([]int{1, -1, -1, 0}))
}
