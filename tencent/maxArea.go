package tencent

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	cur := 0

	for left < right {

		if height[left] <= height[right] {
			val := (right - left) * height[left]
			cur = height[left]

			if val > res {
				res = val
			}

			for left != right {
				left++
				if height[left] > cur {
					break
				}
			}
		} else {
			val := (right - left) * height[right]
			cur = height[right]

			if val > res {
				res = val
			}

			for left != right {
				right--
				if height[right] > cur {
					break
				}
			}
		}
	}

	return res
}
