package tencent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	as := assert.New(t)

	var result float64

	result = FindMedianSortedArrays([]int{}, []int{2})
	as.Equal(2.0, result)

	result = FindMedianSortedArrays([]int{}, []int{2, 3})
	as.Equal(2.5, result)

	result = FindMedianSortedArrays([]int{}, []int{2, 3, 4})
	as.Equal(3.0, result)

	result = FindMedianSortedArrays([]int{1}, []int{2})
	as.Equal(1.5, result)

	result = FindMedianSortedArrays([]int{1, 3}, []int{2})
	as.Equal(2.0, result)

	result = FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	as.Equal(2.5, result)

	result = FindMedianSortedArrays([]int{1, 2, 3}, []int{3, 4})
	as.Equal(3.0, result)

	result = FindMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6})
	as.Equal(3.5, result)
}
