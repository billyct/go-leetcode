package tencent

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	// 第一个数组为空的情况
	if m == 0 {
		return median(nums2)
	}

	// 第二个数组为空的情况
	if n == 0 {
		return median(nums1)
	}

	total := m + n
	totalMiddle := total >> 1 + 1
	// 如果是奇数
	if total & 1 == 1 {
		return findK(nums1, nums2, totalMiddle)
	}

	return (findK(nums1, nums2, totalMiddle - 1) + findK(nums1, nums2, totalMiddle)) / 2
}

func findK(nums1 []int, nums2 []int, k int) float64 {

	m := len(nums1)
	n := len(nums2)

	if m == 0 {
		return float64(nums2[k - 1])
	}

	if n == 0 {
		return float64(nums1[k - 1])
	}

	if k == 1 {
		return math.Min(float64(nums1[0]), float64(nums2[0]))
	}

	km1 := kIndex(k, nums1)
	km2 := kIndex(k, nums2)

	// 如果左边的第 k 数字大于右边的第 k 数字，那么就把右边的第 k 数字前的数字都删掉，k 减少一半
	if nums1[km1 - 1] > nums2[km2 - 1] {
		return findK(nums1, nums2[km2:], k - km2)
	}

	return findK(nums1[km1:], nums2, k - km1)
}

func kIndex(k int, nums []int) int {
	km := k >> 1

	if km > len(nums) {
		return len(nums)
	}

	return km
}

func median(nums []int) float64 {
	n := len(nums)
	k := n >> 1

	// 如果数组长度是奇数那么就取中间的数字
	if n & 1 == 1 {
		return float64(nums[k])
	}

	// 偶数就取中间的两个数字的和除于 2
	return float64(nums[k - 1] + nums[k]) / 2
}
