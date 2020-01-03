package tencent

import "sort"

func ThreeSum(nums []int) [][]int {
	var ret [][]int

	if len(nums) >= 3 {
		sort.Ints(nums)

		for i, m := range nums {
			if m > 0 {
				return ret
			}

			if i > 0 && nums[i] == nums[i - 1] {
				continue
			}

			l := i + 1
			r := len(nums) - 1

			for l < r {
				switch {
				case m+nums[l]+nums[r] == 0:
					ret = append(ret, []int{m, nums[l], nums[r]})

					for l < r && nums[l] == nums[l + 1] {
						l++
					}

					for l < r && nums[r] == nums[r - 1] {
						r--
					}

					l++
					r--
				case m + nums[l] + nums[r] < 0:
					l++
				case m + nums[l] + nums[r] > 0:
					r--
				}
			}
		}
	}

	return ret
}


