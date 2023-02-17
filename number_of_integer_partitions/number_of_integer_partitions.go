package number_of_integer_partitions

func Partitions(n int) int {
	return p(n, n)
}

func p(n int, d int) int {
	if d > n {
		d = n
	}
	if d <= 1 {
		return 1
	}
	count := 0
	for i := d; i > 0; i-- {
		count += p(n-i, i)
	}
	return count
}
