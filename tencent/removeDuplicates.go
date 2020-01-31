package tencent

func RemoveDuplicates(nums []int) int  {
	if len(nums) < 2 {
		return len(nums)
	}

	// 下标
	var i int
	for _, num := range nums[1:] {
		if nums[i] != num {
			i++
			nums[i] = num
		}
	}

	return i + 1
}
