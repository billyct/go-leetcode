package tencent

func merge(A []int, m int, B []int, n int) []int {
	pa, pb, cur := m - 1, n - 1, len(A) - 1

	for pa >= 0 && pb >= 0 {
		if A[pa] < B[pb] {
			A[cur] = B[pb]
			pb--
		} else {
			A[cur] = A[pa]
			pa--
		}

		cur--
	}

	for pb >= 0 {
		A[cur] = B[pb]
		pb--
		cur--
	}

	return A
}

func mergeWithThirdSplice(A []int, m int, B []int, n int)  {
	var pa, pb, cur int
	res := make([]int, m + n)

	for pa < m && pb < n {
		if A[pa] < B[pb] {
			res[cur] = A[pa]
			pa++
		} else {
			res[cur] = B[pb]
			pb++
		}

		cur++
	}

	if pa < m {
		for cur < m + n{
			res[cur] = A[pa]
			cur++
			pa++
		}
	} else {
		for cur < m + n{
			res[cur] = B[pb]
			cur++
			pb++
		}
	}

	copy(A, res)
}
