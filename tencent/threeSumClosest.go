package tencent

import (
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ret := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums) - 2; i++ {
		s, e := i + 1, len(nums) - 1

		for s < e {
			sum := nums[i] + nums[s] + nums[e]

			if abs(target, sum) < abs(target, ret) {
				ret = sum
			}

			if sum < target {
				s++
			} else if sum > target {
				e--
			} else {
				return ret
			}
		}
	}

	return ret
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}