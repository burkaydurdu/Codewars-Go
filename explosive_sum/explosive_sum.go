package explosive_sum

func ExpSum(n uint64) uint64 {
	nums := make([]uint64, n+1)
	nums[0] = 1

	for i := uint64(1); i <= n; i++ {
		for j := i; j <= n; j++ {
			nums[j] += nums[j-i]
		}
	}

	return nums[n]
}
