package find_the_odd_int

// FindOdd - My solution
func FindOdd(seq []int) int {
	var data = make(map[int]int)

	for _, i := range seq {
		data[i]++
	}

	for i, _ := range data {
		if data[i]%2 == 1 {
			return i
		}
	}

	return 0
}

// FindOdd - Clever solution
//func FindOdd(seq []int) int {
//	res := 0
//	for _, x := range seq {
//		res ^= x
//	}
//	return res
//}
