package tencent

func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	res[0] = 1

	for i := 1; i < l; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	tmp := nums[l-1]
	for i := l - 2; i >= 0; i-- {
		res[i] *= tmp
		tmp *= nums[i]
	}

	return res
}
