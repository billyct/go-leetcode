package tencent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	as := assert.New(t)

	as.Equal(2, ThreeSumClosest([]int{-1, 2, 1, -4}, 1))
	as.Equal(0, ThreeSumClosest([]int{0, 2, 1, -3}, 1))
}