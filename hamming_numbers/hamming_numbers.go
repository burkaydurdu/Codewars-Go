package hamming_numbers

func min(x, y uint) uint {
	if x < y {
		return x
	}

	return y
}

func Hammer(n int) uint {
	var list = make([]uint, n)
	var x2, x3, x5 uint = 2, 3, 5
	var i, j, k = 0, 0, 0
	list[0] = 1

	for index := 1; index < n; index++ {
		list[index] = min(x2, min(x3, x5))

		if list[index] == x2 {
			i++
			x2 = 2 * list[i]
		}

		if list[index] == x3 {
			j++
			x3 = 3 * list[j]
		}

		if list[index] == x5 {
			k++
			x5 = 5 * list[k]
		}
	}

	return list[n-1]
}
