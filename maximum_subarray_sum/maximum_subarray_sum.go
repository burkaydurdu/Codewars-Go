package maximum_subarray_sum

func SumOfArray(startIndex, endIndex int, numbers []int) int {
	total := 0
	lengthOfArray := len(numbers)

	if endIndex > lengthOfArray {
		endIndex = lengthOfArray
	}

	if startIndex+1 == endIndex {
		return 0
	}

	numbers = numbers[startIndex:endIndex]

	for _, i := range numbers {
		total += i
	}

	return total
}

func isFullPositiveOrNegative(numbers []int) int {
	var isNegative, isPositive = false, false

	for _, i := range numbers {
		if i < 0 {
			isNegative = true
		} else if i > 0 {
			isPositive = true
		}
	}

	if isPositive && isNegative {
		return 0
	} else if isPositive {
		return 1
	}

	return -1
}

func MaximumSubarraySum(numbers []int) int {
	status := isFullPositiveOrNegative(numbers)
	lengthOfNumber := len(numbers)
	max := 0
	capOfNumber := lengthOfNumber - 2

	if status == 1 {
		return SumOfArray(0, lengthOfNumber, numbers)
	} else if status == -1 {
		return 0
	}

	for i := 0; i < lengthOfNumber; i++ {
		for j := 2; j < capOfNumber; j++ {
			tmp := SumOfArray(i, i+j, numbers)
			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}

// Best Solution
//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
//
//func MaximumSubarraySum(numbers []int) int {
//	res, sum := 0, 0
//	for i := range numbers {
//		sum += numbers[i]
//		res = max(res, sum)
//		sum = max(sum, 0)
//	}
//	return res
//}
