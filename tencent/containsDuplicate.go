package tencent

func containsDuplicate(nums []int) bool {
	tmp := make(map[int]int)

	for _, num := range nums {
		tmp[num]++

		if tmp[num] > 1 {
			return true
		}
	}

	return false
}