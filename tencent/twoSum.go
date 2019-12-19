package tencent

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for v, k := range nums {
		n := target - k
		if _, ok := m[n]; ok {
			return []int{m[n], v}
		}
		m[k] = v
	}

	return []int{}
}
